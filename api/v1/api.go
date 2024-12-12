package api

import (
	"crypto/rsa"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/opengamer-project/og-requests/internal/models"
	"github.com/opengamer-project/og-requests/internal/store"
)

const (
	version = "0.0.1"
	Year    = time.Hour * 8760
)

var (
	// PrivateKey Obviously, this is just a test example. Do not do this in production.
	// In production, you would have the private key and public key pair generated
	// in advance. NEVER add a private key to any GitHub repo.
	PrivateKey *rsa.PrivateKey
)

func apiRootHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "ok",
		"version": version,
	})
}

// InitSecure initializes all secure endpoints for api
func InitSecure(app *fiber.App) {
	app.Get("/api", apiRootHandler)

	app.Get("/logout", func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{Name: "og_auth_token", Value: "", MaxAge: -1})

		return c.Render("logout", nil)
	})

	app.Get("/home", func(c *fiber.Ctx) error {
		return homeGetHandler(c)
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/home", fiber.StatusSeeOther)
	})
}

// Init initializes all insecure endpoints for api
func Init(app *fiber.App) {
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", nil)
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		return loginPOSTHandler(c)
	})

	app.Get("/register", func(c *fiber.Ctx) error {
		return c.Render("register", nil)
	})

	app.Post("/register", func(c *fiber.Ctx) error {
		// Редирект на /home при успешной регистрации
		return registerPOSTHandler(c)
	})

}

func registerPOSTHandler(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user := &models.User{Username: username, Password: password}
	if err := models.CreateUser(store.DB, user); err != nil {
		return c.SendString("Ошибка регистрации: " + err.Error())
	}
	// Create the Claims
	// Create token
	// Generate encoded token and send it as response.
	err := generateJWT(username, c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	c.Response().Header.Set("HX-Redirect", "/home")
	return c.SendString("")
}

func generateJWT(username string, c *fiber.Ctx) error {
	exp := time.Now().Add(Year)

	claims := jwt.MapClaims{
		"name": username,
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

func loginPOSTHandler(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := models.GetUserByUsername(store.DB, username)
	if err != nil || user.Password != password {
		return c.SendString("Неверный логин или пароль")
	}

	err = generateJWT(username, c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Redirect("/home")
}
