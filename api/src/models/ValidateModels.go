package models

import (
	BFE "breakfast/_internal/errors"
	"errors"
	"fmt"
	"reflect"
  "strings"
  "time"
)

// IsModelValid checks if all fields of a struct are populated based on their type.
// It ensures that string fields are non-empty and integer fields are non-zero.
// 
// Type Parameters:
//   - T: The type of the struct being validated.
// 
// Parameters:
//   - s: The struct instance to be validated.
//   - uncheckedFields: A map of field names to be ignored in validation (if the field name is present in this map, it is skipped).
// 
// Returns:
//   - error: Returns an error if any required field is empty or zero, or if the input is not a struct. 
//            Otherwise, it returns nil.
func IsModelValid[T any](s T, uncheckedFields map[string]bool) error {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	if v.Kind() != reflect.Struct {
		return BFE.New(BFE.ErrServer, errors.New("Expected type to be struct"))
	}

	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		if uncheckedFields[fieldName] {
			continue
		}

		field := v.Field(i)
		switch field.Kind() {
		case reflect.String:
			if field.String() == "" {
				return BFE.New(BFE.ErrMissingFields, fmt.Errorf("%v is empty", fieldName))
			}
		case reflect.Int:
			if field.Int() == 0 {
				return BFE.New(BFE.ErrMissingFields, fmt.Errorf("%v is empty", fieldName))
			}
		}
	}
	return nil
}

type ValidationConfig struct {
    IgnoreFields    map[string]bool
    ForbiddenFields map[string]bool
}

// ValidateModel validates the fields of a struct based on a set of allowed and forbidden fields.
// It checks for the presence of unknown fields, forbidden fields, and non-empty required fields.
// Fields that should be ignored or forbidden in validation are provided in the config.
// 
// Type Parameters:
//   - T: The type of the struct being validated.
// 
// Parameters:
//   - s: The struct instance to validate.
//   - requestFields: A map indicating which fields are present in the request. Only these fields will be validated.
//   - config: ValidationConfig with fields that should be ignored or forbidden during validation.
// 
// Returns:
//   - error: Returns an error if the struct contains any unknown or forbidden fields, or if a required field is empty.
//            Returns nil if all fields are valid according to the configuration.
func ValidateModel[T any](s T, requestFields map[string]bool, config ValidationConfig) error {
    v := reflect.ValueOf(s)
    t := reflect.TypeOf(s)

    if v.Kind() != reflect.Struct {
        return BFE.New(BFE.ErrServer, errors.New("expected type to be struct"))
    }

    validFields := make(map[string]bool)
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        jsonTag := field.Tag.Get("json")
        if jsonTag == "" {
            continue
        }

        jsonName := jsonTag
        if comma := strings.Index(jsonTag, ","); comma != -1 {
            jsonName = jsonTag[:comma]
        }
        if jsonName == "-" {
            continue
        }
        validFields[jsonName] = true
    }

    var unknownFields []string
    for fieldName := range requestFields {
        if !validFields[fieldName] {
            unknownFields = append(unknownFields, fieldName)
        }
    }
    if len(unknownFields) > 0 {
        return BFE.New(BFE.ErrUnprocessable, 
            fmt.Errorf("request contains unknown or invalid fields: %s", 
                strings.Join(unknownFields, ", ")))
    }

    for fieldName := range requestFields {
        if config.ForbiddenFields[fieldName] {
            return BFE.New(BFE.ErrUnprocessable, 
                fmt.Errorf("field '%s' cannot be modified in this request", fieldName))
        }
    }

    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        jsonTag := field.Tag.Get("json")
        if jsonTag == "" {
            continue
        }

        jsonName := jsonTag
        if comma := strings.Index(jsonTag, ","); comma != -1 {
            jsonName = jsonTag[:comma]
        }
        if jsonName == "-" {
            continue
        }

        if config.IgnoreFields[field.Name] {
            continue
        }

        if requestFields[jsonName] {
            fieldValue := v.Field(i)
            if isEmpty(fieldValue) {
                return BFE.New(BFE.ErrMissingFields, 
                    fmt.Errorf("field '%s' cannot be empty", jsonName))
            }
        }
    }

    return nil
}

func isEmpty(v reflect.Value) bool {
    switch v.Kind() {
    case reflect.String:
        return v.String() == ""
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        return v.Int() == 0
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
        return v.Uint() == 0
    case reflect.Float32, reflect.Float64:
        return v.Float() == 0
    case reflect.Ptr, reflect.Interface:
        return v.IsNil()
    case reflect.Slice, reflect.Map:
        return v.Len() == 0
    case reflect.Struct:
        if v.Type().String() == "time.Time" {
            return v.Interface().(time.Time).IsZero()
        }
        return false
    default:
        return false
    }
}
