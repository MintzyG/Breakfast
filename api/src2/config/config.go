package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
  DSN string
  JWTSecret string
}

var dsn string
var jwtsecret string

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
    log.Fatalf("Unable to load config!")
	}
	dsn = os.Getenv("MYSQL_DSN")
  if dsn == "" {
    log.Fatalf("No DSN found!")
  }
  jwtsecret = os.Getenv("JWT_KEY")
  if jwtsecret == "" {
    log.Fatalf("No JWT Secret found!")
  }
	return &Config{DSN: dsn, JWTSecret: jwtsecret}
}

func GetJWTSecret() string {
  return jwtsecret
}

func GetDSN() string {
  return dsn
}
