package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	
	server := fiber.New()
	
	
	
	server.Listen(":8080")
}