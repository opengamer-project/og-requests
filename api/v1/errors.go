package api

import "github.com/gofiber/fiber/v2"

func OnJWTError(c *fiber.Ctx, _ error) error {
	return c.Next()
}
