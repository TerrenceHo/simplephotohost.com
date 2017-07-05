package models

import "github.com/jinzhu/gorm"

// Gallery is our image container that viewers view
type Gallery struct {
	gorm.Model
	UserID uint   `gorm:"not null;index"`
	Title  string `gorm:"not null"`
}
