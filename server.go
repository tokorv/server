package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/amber"
	"tokorv.com/apps/TheBookClub"
)

func main() {
	engine := amber.New("./views", ".amber")

	server := fiber.New(fiber.Config{
		Views: engine,
	})

	server.Route("/TheBookClub", func(tbc fiber.Router) {
		tbc.Get("/", func(ctx *fiber.Ctx) error {
			return ctx.SendString("The Book Club!")
		})
	})
	
	server.Listen(":8080")
}