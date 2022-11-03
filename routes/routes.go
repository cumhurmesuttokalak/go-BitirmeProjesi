package routes

import (
	"github.com/cumhurmesuttokalak/go-bitirmeprojesi/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	//user endpoints

	app.Get("/api/getusers", handlers.GetUsers)
	app.Get("/api/getuser/:id", handlers.GetUser)
	app.Post("/api/createuser", handlers.CreateUser)
	app.Delete("/api/deleteuser/:id", handlers.DeleteUser)
	app.Put("/api/updateuser/:id", handlers.UpdateUser)

}
