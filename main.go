package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kravchenkoeag/back_nkah_ecity/db"
	"github.com/kravchenkoeag/back_nkah_ecity/handlers"
)

func main() {
	r := gin.Default()

	// Подключение к базе данных
	db.ConnectToMongoDB()

	// Эндпоинты
	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", handlers.LoginUser)
	r.GET("/user/:id", handlers.GetUserProfile)
	r.PUT("/user/:id", handlers.UpdateUserProfile)

	r.Run()
}
