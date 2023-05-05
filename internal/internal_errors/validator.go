package internalerrors

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(obj any) error {
	validate := validator.New()

	err := validate.Struct(obj)
	if err == nil {
		return nil
	}
	first_error := err.(validator.ValidationErrors)[0]

	return errors.New(first_error.Error())

}
