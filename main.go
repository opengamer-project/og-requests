package main

import (
	"crypto/rand"
	"crypto/rsa"
	"log"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"

	"github.com/opengamer-project/og-requests/api/v1"
	"github.com/opengamer-project/og-requests/internal/store"
)

func main() {
	store.Setup()
	engine := html.New("./templates", ".html")

	// Just as a demo, generate a new private/public key pair on each run. See note above.
	rng := rand.Reader
	var err error
	api.PrivateKey, err = rsa.GenerateKey(rng, 2048)
	if err != nil {
		log.Fatalf("rsa.GenerateKey: %v", err)
	}

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(recover.New())
	api.Init(app)
	app.Static("/static", "./static")

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		TokenLookup: "cookie:og_auth_token",
		ContextKey:  "auth",
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.RS256,
			Key:    api.PrivateKey.Public(),
		},
	}))

	//  ROUTES
	api.InitSecure(app)
	app.Listen(":3000")
}
