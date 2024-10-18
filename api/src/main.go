package main

import (
    "log"
    "fmt"
    "net/http"
    _ "github.com/lib/pq"
    "github.com/joho/godotenv"

    DB "breakfast/repositories"
	UserController "breakfast/controllers/users"
)

func main() {
    if err := start(); err != nil {
        log.Fatal("Error during startup: ", err)
    }
    defer stop()

    mux := http.NewServeMux()
	UserController.Run(mux)

    log.Println("Server starting on :8080")
    if err := http.ListenAndServe(":8080", mux); err != nil {
        log.Fatal("Server error: ", err)
    }
}

func start() error {
    err := godotenv.Load(".env")
    if err != nil {
        return fmt.Errorf("error loading .env: %w", err)
    }

    if err := DB.OpenDatabase(); err != nil {
        return fmt.Errorf("error opening database: %w", err)
    }

    return nil
}

func stop() {
    if err := DB.CloseDatabase(); err != nil {
        log.Printf("Error closing the DB: %v", err)
    }
}
