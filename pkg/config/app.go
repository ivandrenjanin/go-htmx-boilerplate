package config

import "os"

type AppConfig struct {
	Env  string
	Port string
}

func loadAppConfig() AppConfig {
	return AppConfig{
		Env:  os.Getenv("ENV"),
		Port: os.Getenv("PORT"),
	}
}
