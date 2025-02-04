package models

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateModel(model interface{}) (map[string]string, error) {
	if err := validate.Struct(model); err != nil {
		validationErrors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors[err.Field()] = err.Tag()
		}
		return validationErrors, err
	}
	return nil, nil
}
