package db

import (
	"breakfast/internal/models"
	"log"
)

func Migrate() {
	log.Println("Running database migrations...")

	err := DB.AutoMigrate(
		&models.User{},
		&models.UserLink{},
		&models.Maple{},
		&models.MapleDay{},
		&models.Toast{},
		&models.Yogurt{},
		&models.Pancake{},
		&models.CerealActivity{},
		&models.CerealDay{},
		&models.OmeletteCard{},
		&models.OmeletteList{},
		&models.OmeletteTable{},
		&models.ParfaitEvent{},
		&models.ParfaitReminder{},
		&models.EspressoUserSettings{},
		&models.EspressoSession{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Database migrations completed successfully.")
}
