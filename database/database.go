package database

import (
	"todo-app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {

	dsn := "host=localhost user=todo password=password dbname=tasks port=5432 sslmode=disable"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})

}

func Initialize() error {

	db, err := ConnectDB()
	if err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.Task{}); err != nil {
		return err
	}

	return nil

}
