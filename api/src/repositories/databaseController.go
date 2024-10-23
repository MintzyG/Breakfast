package repositories

import (
	BFE "breakfast/errors"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strings"
)

var Instance *sql.DB

func OpenDatabase() error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	Instance, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}

	if err = Instance.Ping(); err != nil {
		return fmt.Errorf("error pinging database: %w", err)
	}

	return nil
}

func CloseDatabase() error {
	if Instance == nil {
		return fmt.Errorf("nil Instance")
	}
	return Instance.Close()
}

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
