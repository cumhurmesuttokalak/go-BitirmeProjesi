package main

import (
	"github.com/cumhurmesuttokalak/go-bitirmeprojesi/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DatabaseConnection()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON("Hello world!")
	})

	app.Listen(":1700")
}
