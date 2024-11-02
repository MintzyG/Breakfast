package repositories

import (
    "fmt"
    "reflect"
    "strings"
)

func BuildUpdateQuery[T any](table string, model T, updates map[string]bool, whereClause string, whereArgs ...interface{}) (string, []interface{}, error) {
    modelValue := reflect.ValueOf(model)
    modelType := reflect.TypeOf(model)
    if modelValue.Kind() == reflect.Ptr {
        modelValue = modelValue.Elem()
        modelType = modelType.Elem()
    }

    fieldMap := make(map[string]int)
    for i := 0; i < modelType.NumField(); i++ {
        field := modelType.Field(i)
        jsonTag := field.Tag.Get("json")
        if jsonTag == "" {
            continue
        }
        if comma := strings.Index(jsonTag, ","); comma != -1 {
            jsonTag = jsonTag[:comma]
        }
        if jsonTag == "-" {
            continue
        }
        fieldMap[jsonTag] = i
    }

    setParts := make([]string, 0, len(updates))
    args := make([]interface{}, 0, len(updates)+len(whereArgs))
    argPosition := 1

    for jsonName := range updates {
        fieldIndex, exists := fieldMap[jsonName]
        if !exists {
            return "", nil, fmt.Errorf("field %s not found in model", jsonName)
        }

        setParts = append(setParts, fmt.Sprintf("%s = $%d", jsonName, argPosition))
        args = append(args, modelValue.Field(fieldIndex).Interface())
        argPosition++
    }

    whereParts := strings.Fields(whereClause)
    newWhereClause := make([]string, len(whereParts))
    whereArgsIndex := argPosition

    for i, part := range whereParts {
        if strings.HasPrefix(part, "$") {
            newWhereClause[i] = fmt.Sprintf("$%d", whereArgsIndex)
            whereArgsIndex++
        } else {
            newWhereClause[i] = part
        }
    }

    whereClause = strings.Join(newWhereClause, " ")
    args = append(args, whereArgs...)

    query := fmt.Sprintf(
        "UPDATE %s SET %s WHERE %s",
        table,
        strings.Join(setParts, ", "),
        whereClause,
    )

    return query, args, nil
}

