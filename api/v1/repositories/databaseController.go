package repositories

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
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

func BeginTransaction() (*sql.Tx, error) {
	return Instance.Begin()
}
