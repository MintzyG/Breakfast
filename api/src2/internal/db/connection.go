package db

import (
  "log"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(dsn string) *gorm.DB {
  var err error
  if dsn == "" {
    log.Fatalf("DSN was empty!")
  }
  DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info),
  })
  if err != nil {
    log.Fatalf("Failed to connect to database: %v", err)
  }
  log.Println("Connected to the database successfully.")
  return DB
}

