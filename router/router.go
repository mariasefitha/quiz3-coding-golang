package router

import (
	"quiz3-coding-golang/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	// CATEGORY ROUTES
	r.POST("/categories", controller.CreateCategory)
	// api.GET("/categories", controller.GetCategories)
	// api.GET("/categories/:id", controller.GetCategoryByID)
	r.GET("/categories", controller.GetAllCategories)
	//api.DELETE("/categories/:id", controller.DeleteCategory)

	//KATEGORI BUKU
	r.POST("/books", controller.CreateBook)
	r.GET("/books", controller.GetAllBooks)
	r.GET("/books/:id", controller.GetBookByID)
	r.PUT("/books", controller.UpdateBook)
	r.DELETE("/books", controller.DeleteBook)

	//KATEGORI + BUKU
	r.GET("/categories/:id/books", controller.GetBooksByCategory)

}
