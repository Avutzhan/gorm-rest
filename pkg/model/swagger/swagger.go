package swagger

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	// docs are generated by Swag CLI, you have to import them.
)

func App() *fiber.App {
	app := fiber.New()
	app.Get("/swagger/*", swagger.Handler) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
	}))
	return app
}