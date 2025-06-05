package models

import (
	"gorm.io/gorm"
)

// User структура для хранения пользователя
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
}
