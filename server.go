package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/amber"
	"fmt"
)

func main() {
	engine := amber.New("./views", ".amber")

	server := fiber.New(fiber.Config{ 
		Views: engine,
	})

	server.Static("/static", "./public")

	server.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World!")
	})
	
	server.Listen(":8080")
}