// Package store handles models storage
package store

import (
	"github.com/gofiber/session/v2"
	"github.com/opengamer-project/og-requests/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Store = session.New()

func Setup() {
	var err error
	DB, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	DB.AutoMigrate(&models.User{})
}
