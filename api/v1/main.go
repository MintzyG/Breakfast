package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	CategoryController "breakfast/controllers/categories"
	PancakeController "breakfast/controllers/pancake"
	ToastController "breakfast/controllers/toast"
	UserController "breakfast/controllers/users"
	YogurtController "breakfast/controllers/yogurt"
	DB "breakfast/repositories"
)

func main() {
	if err := start(); err != nil {
		log.Fatal("Error during startup: ", err)
	}
	defer stop()

	mux := http.NewServeMux()
	UserController.Run(mux)
	CategoryController.Run(mux)
	YogurtController.Run(mux)
	ToastController.Run(mux)
	PancakeController.Run(mux)

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
