package models

import (
	"fmt"
	"reflect"
)

func IsModelValid[T any](s T, excludeFields map[string]bool) error {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	if v.Kind() != reflect.Struct {
    return fmt.Errorf("Expected type to be struct")
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
				return fmt.Errorf("%v is empty", fieldName)
			}
		case reflect.Int:
			if field.Int() == 0 {
				return fmt.Errorf("%v is empty", fieldName)
			}
		}
	}
	return nil
}
