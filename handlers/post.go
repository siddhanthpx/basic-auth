package handlers

import (
	"todo-app/database"
	"todo-app/models"

	"github.com/gofiber/fiber/v2"
)

func AddTask(ctx *fiber.Ctx) error {

	// Parse the JSON object into empty task struct
	var task models.Task
	if err := ctx.BodyParser(&task); err != nil {
		return err
	}

	// Validate the posted JSON
	if err := task.Validate(); err != nil {
		ctx.JSON(fiber.Map{
			"message": err,
		})
		return err
	}

	// Insert the object into the database
	db, err := database.ConnectDB()
	if err != nil {
		ctx.JSON(fiber.Map{
			"message": err,
		})
		return err
	}

	if err := db.Create(&task).Error; err != nil {
		ctx.JSON(fiber.Map{
			"message": err,
		})
		return err
	}

	ctx.JSON(task)

	return nil
}
