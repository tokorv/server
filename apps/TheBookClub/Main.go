package TheBookClub

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/amber"
)

func Init() *fiber.App {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello there!")
	})

	return app
}