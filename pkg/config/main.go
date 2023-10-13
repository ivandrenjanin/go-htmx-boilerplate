package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
}

func NewConfig() *Config {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file \n %v", err)
	}

	return &Config{
		App:      loadAppConfig(),
		Database: loadDatabaseConfig(),
	}
}
