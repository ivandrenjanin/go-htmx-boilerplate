package config

import "os"

type HTTPConfig struct {
	Host string
	Port string
	Env  string
}

func LoadHTTPConfig() HTTPConfig {
	return HTTPConfig{
		Host: os.Getenv("APP_HOST"),
		Port: os.Getenv("APP_PORT"),
		Env:  os.Getenv("APP_ENV"),
	}
}
