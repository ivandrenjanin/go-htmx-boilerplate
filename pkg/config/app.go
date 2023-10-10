package config

import "os"

type AppConfig struct {
	Host          string
	Env           string
	Port          string
	AccessSecret  string
	RefreshSecret string
	CSRFSecret    string
}

func loadAppConfig() AppConfig {
	return AppConfig{
		Host:          os.Getenv("APP_HOST"),
		Env:           os.Getenv("ENV"),
		Port:          os.Getenv("PORT"),
		AccessSecret:  os.Getenv("ACCESS_SECRET"),
		RefreshSecret: os.Getenv("REFRESH_SECRET"),
		CSRFSecret:    os.Getenv("CSRF_SECRET"),
	}
}
