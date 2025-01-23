package db

import (
	"breakfast/internal/models"
	"log"
)

func Migrate() {
	log.Println("Running database migrations...")

	err := DB.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Database migrations completed successfully.")
}
