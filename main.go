package main

import (
	"github.com/gin-gonic/gin"
	"github.com/username/back_nkah_ecity/db"
	"github.com/username/back_nkah_ecity/handlers"
)

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Connect to MongoDB
	db.ConnectToMongoDB()

	// Define routes (endpoints)
	router.POST("/register", handlers.RegisterUser)
	router.POST("/login", handlers.LoginUser)
	router.GET("/user/:id", handlers.GetUserProfile)
	router.PUT("/user/:id", handlers.UpdateUserProfile)
	router.GET("/groups", handlers.GetGroups)
	router.POST("/groups", handlers.CreateGroup)

	// Start the server on port 8080
	router.Run(":8080")
}
