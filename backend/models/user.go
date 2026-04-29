package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required,email" gorm:"unique;not null"`
	Password string `json:"password,omitempty" binding:"required,min=6" gorm:"not null"`
}