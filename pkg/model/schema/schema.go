package schema

import (
	"github.com/gofiber/fiber/v2"
	"gorm-rest/pkg/model/booking"
)

type Schema struct {
	Group       string        `json:"group"`
	UrlName     string        `json:"url_name"`
	ChildSchema []ChildSchema `json:"child_schema"`
	Schema      []interface{} `json:"schema"`
	Order       int           `json:"order"`
}
type ChildSchema struct {
	Group   string       `json:"group"`
	UrlName string       `json:"url_name"`
	Schema  []SchemaMenu `json:"schema"`
	Order   int          `json:"order"`
	Icon    string       `json:"icon,omitempty"`
}
type SchemaMenu struct {
	Name  string `json:"name"`
	Order int    `json:"order"`
	Icon  string `json:"icon,omitempty"`
}

func App() *fiber.App {
	mainSchema := Schema{
		Group:   "main",
		UrlName: "",
		ChildSchema: []ChildSchema{
			{
				Group:   "Booking",
				UrlName: "booking",
				Order:   0,
				Icon:    "face",
				Schema: []SchemaMenu{
					{
						Name:  "new-booking",
						Order: 0,
						Icon:  "",
					},
					{
						Name:  "all-booking",
						Order: 0,
						Icon:  "",
					},
					{
						Name:  "booking",
						Order: 0,
						Icon:  "",
					},
					{
						Name:  "update-booking",
						Order: 0,
						Icon:  "",
					},
				},
			},
		},
		Schema: nil,
		Order:  0,
	}

	app := fiber.New()

	app.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(mainSchema)
	})
	app.Mount("", booking.SchemaApp())
	return app
}
