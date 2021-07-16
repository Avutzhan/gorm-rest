package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"gorm-rest/pkg/ftools/fconf"
	"gorm-rest/pkg/ftools/fdb"
	"gorm-rest/pkg/ftools/flog"
	"gorm-rest/pkg/ftools/fmigrate"
	"gorm-rest/pkg/model/booking"
	"gorm-rest/pkg/model/history"
	"gorm-rest/pkg/model/schema"
	"gorm-rest/pkg/model/swagger"
)

func Start() {
	if result := fdb.Do(); result != nil {
		flog.Error().Err(result).Msg("error connect to database. app close")

		return
	}
	defer fdb.Close()

	if fconf.DatabaseMigration() {
		fmigrate.Run(
			booking.Booking{},
		)
	}

	run()
}
func run() {
	app := fiber.New()
	if fconf.UseCors() {
		app.Use(cors.New(cors.Config{
			AllowOrigins:     "http://192.168.2.9:8085,http://192.168.2.9:8085/,http://192.168.1.36:8085/,http://192.168.1.36:8085",
			AllowMethods:     "",
			AllowHeaders:     "content-type,x-model,x-sleep,x-full",
			AllowCredentials: true,
			ExposeHeaders:    "content-type,x-model,x-sleep,x-full",
			MaxAge:           0,
		}))
	}
	app.Use(func(c *fiber.Ctx) error {
		defer func() {
			if err := recover(); err != nil {
				flog.ErrorCtx(c).Interface("err", err).Msg(">>>>>")

				if err := c.SendStatus(fiber.StatusInternalServerError); err != nil {
					flog.ErrorCtx(c).Err(err).Msg(">>>>")
				}
			}
		}()
		return c.Next()
	})
	app.Use(pprof.New())
	app.Use(history.New)

	app.Mount("", swagger.App())
	api := app.Group("/v1")
	football := api.Group("/football")
	football.Mount("/booking", booking.App())

	doc := api.Group("/schema")
	doc.Mount("", schema.App())

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	})
	flog.Error().Err(app.Listen(fconf.ServerUP())).Msg("failed run server")
}
