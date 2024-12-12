package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User структура для хранения пользователя
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
}

// Connect подключается к базе данных SQLite, используя локальный файл
func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("local.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Миграция модели
	db.AutoMigrate(&User{})
	return db, nil
}

// CreateUser создает нового пользователя
func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

// GetUserByUsername ищет пользователя по имени
func GetUserByUsername(db *gorm.DB, username string) (*User, error) {
	var user User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
