package routes

import (
	"github.com/gofiber/fiber/v2"
	"goApiTask/handlers"
)

func RegisterTaskRoutes(app *fiber.App, client *handlers.Client) {
	api := app.Group("/api")
	api.Get("/", client.GetPage)
	api.Post("/create", client.Create)
	api.Get("/getalltasks", client.GetTasks)
	api.Put("/update/:id", client.Update)
	api.Delete("/delete/:id", client.Delete)
}
