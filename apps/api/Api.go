package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/amber"
)

func Init() fiber.App {
	engine := amber.New("./views", ".amber")

	app := fiber.New(fiber.Config{
		View: engine,
	})
}

