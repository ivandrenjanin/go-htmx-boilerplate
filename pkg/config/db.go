package config

import "os"

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	Sslmode  string
	Driver   string
}

func loadDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Sslmode:  os.Getenv("DB_SSL_MODE"),
		Driver:   os.Getenv("DB_DRIVER"),
	}
}
