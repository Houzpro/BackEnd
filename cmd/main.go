package main

import (
	"gin-notes-api/internal/db"
	"gin-notes-api/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8080")
}
