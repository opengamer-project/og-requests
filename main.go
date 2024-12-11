package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
	"github.com/opengamer-project/og-requests/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var store = session.New()

func setupDatabase() {
	var err error
	db, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&models.User{})
}

func main() {
	setupDatabase()

	app := fiber.New()

	app.Static("/static", "./static")

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("./templates/login.html", nil)
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		user, err := models.GetUserByUsername(db, username)
		if err != nil || user.Password != password {
			return c.SendString("Неверный логин или пароль")
		}

		// Успешный вход
		session := store.Get(c)
		session.Set("user", user.Username)
		session.Save()

		return c.Redirect("/home")
	})

	app.Get("/register", func(c *fiber.Ctx) error {
		return c.Render("./templates/register.html", nil)
	})

	app.Post("/register", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		user := &models.User{Username: username, Password: password}
		if err := models.CreateUser(db, user); err != nil {
			return c.SendString("Ошибка регистрации: " + err.Error())
		}

		return c.SendString("Регистрация успешна!")
	})

	app.Get("/home", func(c *fiber.Ctx) error {
		session := store.Get(c)
		username := session.Get("user")
		return c.SendString("Добро пожаловать, " + username.(string))
	})

	app.Listen(":3000")
}
