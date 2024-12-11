package api

import "github.com/gofiber/fiber/v2"

const version = "0.0.1"

func apiRootHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "ok",
		"version": version,
	})
}

// Init initializes all endpoints for api
func Init(app *fiber.App, version string) {
	app.Get("/", apiRootHandler)
	api := app.Group("/" + version)
	api.Get("/", apiRootHandler)

}
