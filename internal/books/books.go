package books

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Year      int    `json:"year"`
	Publisher string `json:"publisher,omitempty"`
}

func SetDB(database *gorm.DB) {
	db = database
}

func GetBooks(c *gin.Context) {
	var books []Book
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "3"))
	offset := (page - 1) * pageSize

	if db == nil {
		log.Println("Database connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection is nil"})
		return
	}

	query := db.Limit(pageSize).Offset(offset)

	if title := c.Query("title"); title != "" {
		query = query.Where("title ILIKE ?", "%"+title+"%")
	}
	if author := c.Query("author"); author != "" {
		query = query.Where("author ILIKE ?", "%"+author+"%")
	}

	if sort := c.Query("sort"); sort != "" {
		query = query.Order(sort)
	}

	if err := query.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch books"})
		return
	}

	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	var book Book
	if err := db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	if err := db.Create(&newBook).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create book"})
		return
	}
	c.JSON(http.StatusCreated, newBook)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var updatedBook Book
	if err := c.BindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	if err := db.Model(&Book{}).Where("id = ?", id).Updates(updatedBook).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	c.JSON(http.StatusOK, updatedBook)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&Book{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "book deleted"})
}

func GetBooksByYearRange(c *gin.Context) {
	startYear := c.Query("startYear")
	endYear := c.Query("endYear")

	var books []Book
	if err := db.Where("year BETWEEN ? AND ?", startYear, endYear).Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching books"})
		return
	}
	c.JSON(http.StatusOK, books)
}

func UpdatePublishers(c *gin.Context) {
	newPublisher := c.Query("publisher")
	author := c.Query("author")
	if newPublisher == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "publisher parameter is required"})
		return
	}
	if author == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "author parameter is required"})
		return
	}

	log.Println("Starting transaction to update publishers")

	tx := db.Begin()
	if tx.Error != nil {
		log.Println("Failed to start transaction")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to start transaction"})
		return
	}
	log.Println("Transaction started")

	if err := tx.Model(&Book{}).Where("author = ?", author).Update("publisher", newPublisher).Error; err != nil {
		log.Println("Failed to update publishers, rolling back transaction")
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update publishers"})
		return
	}
	log.Println("Publishers updated successfully")

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed to commit transaction, rolling back")
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to commit transaction"})
		return
	}
	log.Println("Transaction committed successfully")

	c.JSON(http.StatusOK, gin.H{"message": "publishers updated successfully"})
}

func CountBooksByAuthor(c *gin.Context) {
	var results []struct {
		Author string
		Count  int
	}

	if err := db.Model(&Book{}).
		Select("author, COUNT(*) as count").
		Group("author").
		Scan(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to count books by author"})
		return
	}

	c.JSON(http.StatusOK, results)
}
