package routes

import (
	"gin-notes-api/internal/auth"
	"gin-notes-api/internal/books"
	"gin-notes-api/internal/tasks"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/login", auth.Login)
	router.POST("/register", auth.Register)
	router.POST("/refresh", auth.Refresh)

	protected := router.Group("/")
	protected.Use(auth.AuthMiddleware())
	{
		protected.GET("/books", books.GetBooks)
		protected.GET("/books/:id", books.GetBookByID)
		protected.POST("/books", auth.RoleMiddleware("admin"), books.CreateBook)
		protected.PUT("/books/:id", auth.RoleMiddleware("admin"), books.UpdateBook)
		protected.DELETE("/books/:id", auth.RoleMiddleware("admin"), books.DeleteBook)
		protected.GET("/booksByYearRange", books.GetBooksByYearRange)
		protected.PUT("/updatePublishers", books.UpdatePublishers)
		protected.GET("/countBooksByAuthor", books.CountBooksByAuthor)
	}

	// Task routes
	router.POST("/tasks", tasks.CreateTaskHandler)
	router.GET("/tasks/:id", tasks.GetTaskHandler)
	router.DELETE("/tasks/:id", tasks.CancelTaskHandler)
}
