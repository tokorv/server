package TheBookClub

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/amber"
)

func Init() *fiber.App {
	engine := amber.New("./views", ".amber")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("/TheBookClub/index", fiber.Map {
			"Title": "The Book Club",
		})
	})

	return app
}