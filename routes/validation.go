package routes

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func validateStruct(todo interface{}) error {
	return validate.Struct(todo)
}
