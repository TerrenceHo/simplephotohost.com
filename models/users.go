package models

import (
	"github.com/jinzhu/gorm"
	_ "github.github.com/jinzhu/gorm/dialect/postgres"
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}
