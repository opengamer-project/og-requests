package api

import (
	"html/template"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/opengamer-project/og-requests/internal/models"
)

func initRequests(app *fiber.App) {
	app.Get("/requests", requestsGetHandler)
}

func requestsGetHandler(c *fiber.Ctx) error {
	name := ""

	userBefore := c.Locals("auth")
	if userBefore == nil {
		return c.Redirect("/login") // TODO: return to
	}

	claims := userBefore.(*jwt.Token).Claims.(jwt.MapClaims)
	name = claims["name"].(string)

	tmp := []models.Claim{models.Claim{}, models.Claim{}}
	renderedClaims := []template.HTML{}

	for _, requsest := range tmp {
		t, err := requsest.Render()
		if err != nil {
			return err
		}
		renderedClaims = append(renderedClaims, t)
	}

	return c.Render("requests", fiber.Map{
		"Title":    "Заявки",
		"Name":     name,
		"Requests": renderedClaims,
	})

}
