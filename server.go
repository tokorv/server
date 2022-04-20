package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/amber"
)

func main() {
	
	engine := amber.New("./views", ".amber")

	server := fiber.New(fiber.Config{
		Views: engine,
	})

	server.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("index", fiber.Map {
			"Title": "Will likes poo",
		})
	})
	
	server.Listen(":8080")
}