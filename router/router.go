package router

import (
	"quiz3-coding-golang/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	api := r.Group("/api")
	{
		// CATEGORY ROUTES
		api.POST("/categories", controller.CreateCategory)
		// api.GET("/categories", controller.GetCategories)
		// api.GET("/categories/:id", controller.GetCategoryByID)
		api.GET("/categories", controller.GetAllCategories)
		//api.DELETE("/categories/:id", controller.DeleteCategory)

		//KATEGORI BUKU
		api.POST("/books", controller.CreateBook)
		api.GET("/books", controller.GetAllBooks)
		api.GET("/books/:id", controller.GetBookByID)
		api.PUT("/books", controller.UpdateBook)
		api.DELETE("/books", controller.DeleteBook)

		//KATEGORI + BUKU
		api.GET("/categories/:id/books", controller.GetBooksByCategory)

	}

}
