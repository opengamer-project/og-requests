package api

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/opengamer-project/og-requests/internal/models"
	"github.com/opengamer-project/og-requests/internal/store"
)

func initLogin(app *fiber.App) {
	app.Get("/logout", func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{Name: "og_auth_token", Value: "", MaxAge: -1})

		return c.Render("logout", fiber.Map{
			"Title": "OG Portal",
		})
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{
			"Title": "OG Portal: Логин",
		})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		return loginPOSTHandler(c)
	})

	app.Get("/register", func(c *fiber.Ctx) error {
		return c.Render("register", fiber.Map{
			"Title": "OG Portal: Регистрация",
		})
	})

	app.Post("/register", func(c *fiber.Ctx) error {
		// Редирект на /home при успешной регистрации
		return registerPOSTHandler(c)
	})
}

func loginPOSTHandler(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := store.GetUserByUsername(store.DB, username)
	if err != nil || user.Password != password {
		return c.SendString("Неверный логин или пароль")
	}

	err = generateJWT(username, user.ID, c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Redirect("/home")
}

func registerPOSTHandler(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user := &models.User{Username: username, Password: password}
	if err := store.CreateUser(store.DB, user); err != nil {
		return c.SendString("Ошибка регистрации: " + err.Error())
	}
	// Create the Claims
	// Create token
	// Generate encoded token and send it as response.
	err := generateJWT(username, user.ID, c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	c.Response().Header.Set("HX-Redirect", "/home")
	return c.SendString("")
}

func generateJWT(username string, id uint, c *fiber.Ctx) error {
	exp := time.Now().Add(Year)

	claims := jwt.MapClaims{
		"name": username,
		"id":   id,
		"exp":  exp.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	t, err := token.SignedString(PrivateKey)
	if err != nil {
		return err
	}
	c.Cookie(&fiber.Cookie{Name: "og_auth_token", Value: t, Expires: exp})
	return nil
}
