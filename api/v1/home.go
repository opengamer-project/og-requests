package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

func homeGetHandler(c *fiber.Ctx) error {
	userBefore := c.Locals("auth")
	if userBefore == nil {
		log.Warn(errors.New("cannto get user info from auth"))
		return c.SendString("Доступа нет")
	}
	claims := userBefore.(*jwt.Token).Claims.(jwt.MapClaims)
	name, ok := claims["name"].(string)
	if name != "" && ok {
		return c.Render("home", fiber.Map{
			"Title": "OG Portal",
			"Name":  name,
		})
	}
	return c.SendString("Доступа нет")

}
