package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

// Config contains application configs
type Config struct {
	Environment string
	Port        string
	Database    *Database
}

// Database contains database configs
type Database struct {
	Host     string
	Port     string
	User     string
	DB       string
	Password string
}

// NewConfig returns a struct containing application configuration
func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Error("Cannot load .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		Environment: os.Getenv("ENV"),
		Port:        port,
		Database: &Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			DB:       os.Getenv("DB_DB"),
			Password: os.Getenv("DB_PASSWORD"),
		},
	}, nil
}
