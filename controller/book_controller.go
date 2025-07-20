package controller

import (
	"net/http"
	"quiz3-coding-golang/config"
	"quiz3-coding-golang/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CREATE BOOK
func CreateBook(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan buku"})
		return
	}
	c.JSON(http.StatusCreated, book)

}

// GET ALL BOOKS
func GetAllBooks(c *gin.Context) {
	var books []model.Book
	if err := config.DB.Preload("Category").Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data buku"})
		return
	}
	c.JSON(http.StatusOK, books)
}

// GET BOOK BY ID
func GetBookByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book model.Book
	if err := config.DB.Preload("Category").First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// func GetBooks(c *gin.Context) {
// 	var books []model.Book
// 	if err := config.DB.Preload("Category").Find(&books).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, books)
// }

// UPDATE BOOK
func UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book model.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
		return
	}
	var updateData model.Book
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book.Title = updateData.Title
	book.Author = updateData.Author
	book.CategoryID = updateData.CategoryID

	config.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}

// DETELE BOOK
func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book model.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
		return
	}
	config.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Buku berhasil dihapus~~"})
}

// GET BOOK BY KATEGORI
func GetBooksByCategory(c *gin.Context) {
	categoryID := c.Param("id")

	var books []model.Book
	if err := config.DB.Where("category_id = ?", categoryID).Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal ambil data"})
		return
	}

	c.JSON(http.StatusOK, books)
}
