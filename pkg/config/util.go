package config

import (
	"log"
	"os"
	"strconv"
)

const (
	// APP
	APP_ENV          = "APP_ENV"
	APP_HOST         = "APP_HOST"
	APP_PORT         = "APP_PORT"
	APP_DEBUG        = "APP_DEBUG"
	APP_READ_TIMEOUT = "APP_READ_TIMEOUT"

	// DB
	DB_HOST                     = "DB_HOST"
	DB_PORT                     = "DB_PORT"
	DB_USER                     = "DB_USER"
	DB_PASSWORD                 = "DB_PASSWORD"
	DB_NAME                     = "DB_NAME"
	DB_DEBUG                    = "DB_DEBUG"
	DB_SSL_MODE                 = "DB_SSL_MODE"
	DB_MAX_OPEN_CONNECTIONS     = "DB_MAX_OPEN_CONNECTIONS"
	DB_MAX_IDLE_CONNECTIONS     = "DB_MAX_IDLE_CONNECTIONS"
	DB_MAX_LIFETIME_CONNECTIONS = "DB_MAX_LIFETIME_CONNECTIONS"

	// JWT
	JWT_SECRET_KEY                      = "JWT_SECRET_KEY"
	JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT = "JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"
)

func GetEnvString(key string) string {
	return os.Getenv(key)
}

func GetEnvInt(key string) int {
	v, err := strconv.Atoi(os.Getenv(key))

	if err != nil {
		log.Fatalf("error fetching variable %s, err: %v", key, err)
	}

	return v
}

func GetEnvBool(key string) bool {
	v, err := strconv.ParseBool(os.Getenv(key))

	if err != nil {
		log.Fatalf("error fetching variable %s, err: %v", key, err)
	}

	return v
}
