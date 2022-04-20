package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	server := fiber.New()
	
	server.Get("/", func(ctx *fiber.Ctx) {
		return ctx.SendString("Hello!")
	})
	
	server.Listen(":8080")
}