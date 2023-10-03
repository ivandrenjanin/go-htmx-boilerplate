package config

import (
	"time"
)

type App struct {
	Env         string
	Host        string
	Port        int
	Debug       bool
	ReadTimeout time.Duration

	// JWT Conf
	JWTSecretKey               string
	JWTSecretExpireMinuteCount int
}

var app = new(App)

func AppCfg() *App {
	return app
}

func LoadApp() {
	app.Env = GetEnvString(APP_ENV)
	app.Host = GetEnvString(APP_HOST)
	app.Port = GetEnvInt(APP_PORT)
	app.Debug = GetEnvBool(APP_DEBUG)

	readTimeout := GetEnvInt(APP_READ_TIMEOUT)
	app.ReadTimeout = time.Duration(readTimeout) * time.Second

	app.JWTSecretKey = GetEnvString(JWT_SECRET_KEY)
	app.JWTSecretExpireMinuteCount = GetEnvInt(JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT)
}
