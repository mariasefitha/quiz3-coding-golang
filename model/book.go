package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title" gorm:"type:varchar(100);not null" binding:"required"`
	Author      string `json:"author" gorm:"type:varchar(100);not null" binding:"required"`
	Description string `json:"description" binding:"required"`
	CategoryID  uint   `json:"category_id" binding:"required"`
	// Category   Category `json:"category" gorm:"foreignKey:CategoryID"`
}
