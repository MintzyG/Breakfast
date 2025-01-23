package db

import (
    "log"
    "breakfast/internal/models"
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

