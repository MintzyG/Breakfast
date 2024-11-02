package repositories

import (
	BFE "breakfast/_internal/errors"
	"errors"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

func BuildUpdateQuery(table string, updates map[string]interface{}, validFields map[string]bool, whereClause string, whereArgs ...interface{}) (string, []interface{}, error) {
	if len(updates) == 0 {
		return "", nil, BFE.New(BFE.ErrUnprocessable, errors.New("No fields in request"))
	}

	setParts := []string{}
	args := []interface{}{}
	argIdx := 1

	for key, value := range updates {
		if _, valid := validFields[key]; valid {
			setParts = append(setParts, fmt.Sprintf("%s = $%d", key, argIdx))
			args = append(args, value)
			argIdx++
		} else {
			return "", nil, BFE.New(BFE.ErrUnprocessable, fmt.Errorf("Invalid field: %s", key))
		}
	}

	if len(setParts) == 0 {
		return "", nil, BFE.New(BFE.ErrUnprocessable, errors.New("No valid fields to update"))
	}

	for i := range whereArgs {
		whereClause = strings.Replace(whereClause, fmt.Sprintf("$%d", i+1), fmt.Sprintf("$%d", argIdx), 1)
		args = append(args, whereArgs[i])
		argIdx++
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s", table, strings.Join(setParts, ", "), whereClause)

	return query, args, nil
}
