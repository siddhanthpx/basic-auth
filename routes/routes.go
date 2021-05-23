package routes

import (
	"todo-app/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// Version 1 of the API
	v1 := app.Group("/api/v1")

	v1.Get("/", handlers.GetAllTasks)
	v1.Get("/:id", handlers.GetTask)

	v1.Post("/", handlers.AddTask)
}
