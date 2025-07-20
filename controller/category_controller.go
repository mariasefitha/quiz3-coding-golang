package controller

import (
	"net/http"
	"quiz3-coding-golang/config"
	"quiz3-coding-golang/model"

	"github.com/gin-gonic/gin"
)

// CREATE KATEGORI
func CreateCategory(c *gin.Context) {
	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

// GET KATEGORI
func GetCategories(c *gin.Context) {
	var categories []model.Category
	if err := config.DB.Preload("Books").Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// GET KATEGORI BY ID
func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	var category model.Category
	if err := config.DB.Preload("Books").First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, category)
}

// GET ALL KATEGORI
func GetAllCategories(c *gin.Context) {
	var categories []model.Category
	result := config.DB.Find(&categories)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// DELETE KATEGORI
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	var category model.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori tidak ditemukan"})
		return
	}

	if err := config.DB.Delete(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kategori berhasil dihapus~~"})
}
