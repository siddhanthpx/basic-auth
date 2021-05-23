package handlers

import (
	"todo-app/database"
	"todo-app/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllTasks(ctx *fiber.Ctx) error {

	db, err := database.ConnectDB()
	if err != nil {
		ctx.JSON(fiber.Map{
			"message": err,
		})
		return err
	}

	var tasks []models.Task
	db.Table("tasks").Scan(&tasks)

	return ctx.JSON(&tasks)

}

func GetTask(ctx *fiber.Ctx) error {

	// Parse id from query
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	db, err := database.ConnectDB()
	if err != nil {
		ctx.JSON(fiber.Map{
			"message": err,
		})
		return err
	}

	var task models.Task
	if err := db.Where(&models.Task{ID: id}).First(&task).Error; err != nil {
		return err
	}

	return ctx.JSON(&task)

}
