package models

import (
  BFE "breakfast/errors"
	"fmt"
	"reflect"
)

func IsModelValid[T any](s T, excludeFields map[string]bool) error {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	if v.Kind() != reflect.Struct {
    return BFE.NewBFError(BFE.SERVER_ERROR_CODE, "Expected type to be struct")
	}

	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		if excludeFields[fieldName] {
			continue
		}

		field := v.Field(i)
		switch field.Kind() {
		case reflect.String:
			if field.String() == "" {
        return BFE.NewBFError(BFE.MISSING_FIELDS_CODE, fmt.Sprintf("%v is empty", fieldName))
			}
		case reflect.Int:
			if field.Int() == 0 {
        return BFE.NewBFError(BFE.MISSING_FIELDS_CODE, fmt.Sprintf("%v is empty", fieldName))
			}
		}
	}
	return nil
}
