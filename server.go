package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/amber"
	"tokorv.com/apps/TheBookClub"
)

func main() {
	BookClub := TheBookClub.Init()
	engine := amber.New("./views", ".amber")

	server := fiber.New(fiber.Config{
		Views: engine,
	})

	server.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("index", fiber.Map {
			"Title": "Will likes poo",
		})
	})

	server.Mount("/TheBookClub", BookClub)
	
	server.Listen(":8080")
}