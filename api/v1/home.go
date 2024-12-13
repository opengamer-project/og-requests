package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func homeGetHandler(c *fiber.Ctx) error {
	msg := ""
	name := ""
	isAuthenticated := false

	userBefore := c.Locals("auth")
	if userBefore == nil {
		msg = "Доступа нет"
	} else {
		claims := userBefore.(*jwt.Token).Claims.(jwt.MapClaims)
		name, isAuthenticated = claims["name"].(string)
	}

	return c.Render("home", fiber.Map{
		"Title":         "OG Portal",
		"Name":          name,
		"Msg":           msg,
		"Authenticated": isAuthenticated,
	})

}
