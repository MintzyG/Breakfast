package models

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateModel(model interface{}) (map[string]string, error) {
	if err := validate.Struct(model); err != nil {
		validationErrors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "eq":
				validationErrors[err.Field()] = fmt.Sprintf("%s must be exactly %s", err.Field(), err.Param())
			case "min":
				validationErrors[err.Field()] = fmt.Sprintf("%s must be at least %s", err.Field(), err.Param())
			case "max":
				validationErrors[err.Field()] = fmt.Sprintf("%s must be at most %s", err.Field(), err.Param())
			case "required":
				validationErrors[err.Field()] = fmt.Sprintf("%s is required", err.Field())
			case "email":
				validationErrors[err.Field()] = fmt.Sprintf("%s must be a valid email", err.Field())
			default:
				validationErrors[err.Field()] = fmt.Sprintf("Invalid value for %s", err.Field())
			}
		}
		return validationErrors, err
	}
	return nil, nil
}
