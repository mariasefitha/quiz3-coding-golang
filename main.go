package main

import (
	"log"
	"os"
	"quiz3-coding-golang/config"
	"quiz3-coding-golang/model"
	"quiz3-coding-golang/router"

	"quiz3-coding-golang/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	// Connect to DB & AutoMigrate
	config.ConnectDB()
	config.DB.AutoMigrate(&model.Category{}, &model.Book{})
	config.DB.AutoMigrate(&model.Book{}, &model.Category{})

	//MIDDLEWARE
	protected := r.Group("/api", middleware.BasicAuthMiddleware())

	// Register routes
	router.RegisterRoutes(protected)

	// Root path to verify running status
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API is running!",
		})
	})

	// Run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)

}
