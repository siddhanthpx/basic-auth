package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {

	// Version 1 of the API
	v1 := app.Group("/api/v1")
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"ping": "pong",
		})
	})
}
