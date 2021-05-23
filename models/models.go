package models

import "github.com/go-playground/validator"

// Task Model
type Task struct {
	ID     int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name   string `json:"name" validate:"required"`
	Status bool   `json:"status"`
}

func (g *Task) Validate() error {

	validate := validator.New()
	return validate.Struct(g)

}
