package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model

	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required,max=1000"`
}

var validate = validator.New()

func (t *Todo) Validate() error {
	return validate.Struct(t)
}
