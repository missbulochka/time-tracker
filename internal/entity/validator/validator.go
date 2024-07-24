package validator

import (
	"github.com/go-playground/validator/v10"
)

func ValidateStruct(model struct{}) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(model); err != nil {
		return err
	}

	return nil
}
