package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Name  string `gorm:"type:text;not null"`
	Age   int    `gorm:"not null"`
	Email string `gorm:"type:text;unique;not null"`
}
