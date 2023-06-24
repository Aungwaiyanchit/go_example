package main

import (
	"github.com/Aungwaiyanchit/books/controllers"
	"github.com/Aungwaiyanchit/books/middleware"
	"github.com/Aungwaiyanchit/books/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDb()

	protected := router.Group("/")
	protected.Use(middleware.JwtAuthMiddleware())
	protected.GET("books", controllers.FindBooks)
	protected.POST("books", controllers.CreateBook)
	protected.GET("books/:id", controllers.FindBookById)
	protected.PATCH("books/:id", controllers.UpdateBook)
	protected.DELETE("books/:id", controllers.DeleteBook)


	router.GET("/users", controllers.FindUsers)
	router.POST("/users", controllers.CreatUser)
	router.POST("/login", controllers.LoginUser)

	router.Run("localhost:8080")
}