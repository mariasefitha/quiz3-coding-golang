package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name" gorm:"type:varchar(100);not null"`
	Books []Book `json:"books,omitempty"`
}
