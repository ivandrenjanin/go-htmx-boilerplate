package config

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func LoadAllConfigs(envfile string) {
	err := godotenv.Load(envfile)
	if err != nil {
		log.Fatalf("can not load .env file: %v", err)
	}

	LoadApp()
	LoadDBCfg()
}

func FiberConfig() fiber.Config {
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(AppCfg().ReadTimeout),
	}
}
