package main

import (
	"github.com/cumhurmesuttokalak/go-bitirmeprojesi/database"
	"github.com/cumhurmesuttokalak/go-bitirmeprojesi/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DatabaseConnection()
	app := fiber.New()
	routes.SetupRoutes(app)
	app.Listen(":1700")
}
