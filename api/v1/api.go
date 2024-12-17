package api

import (
	"crypto/rsa"
	"time"

	"github.com/gofiber/fiber/v2"
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

// Init initializes all insecure endpoints for api
func Init(app *fiber.App) {
	app.Get("/api", apiRootHandler)

	initHome(app)
	initRequests(app)
	initLogin(app)
}
