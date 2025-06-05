// Package store handles models storage
package store

import (
	"github.com/gofiber/session/v2"
	"github.com/opengamer-project/og-requests/internal/models"
	"github.com/pkg/errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	Store = session.New()
)

func Create(c *models.Claim) error {
	return DB.Create(*c).Error
}

// Connect подключается к базе данных SQLite, используя локальный файл
func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("local.db"), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "connection to db")
	}

	// Миграция модели
	if err := db.AutoMigrate(&models.User{}, &models.Claim{}); err != nil {
		return nil, errors.Wrap(err, "aplying migration")
	}
	return db, nil
}

// CreateUser создает нового пользователя
func CreateUser(db *gorm.DB, user *models.User) error {
	return db.Create(user).Error
}

// GetUserByUsername ищет пользователя по имени
func GetUserByUsername(db *gorm.DB, username string) (*models.User, error) {
	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func Setup() {
	var err error
	DB, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
}
