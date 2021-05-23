package main

import (
	"todo-app/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":8080")

}
