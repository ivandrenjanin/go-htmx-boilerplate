package logger

import (
	"os"

	"github.com/ivandrenjanin/go-fiber-htmx-boilerplate/pkg/config"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

var logger = new(Logger)

func SetUpLogger() {
	logger = &Logger{logrus.New()}

	appEnv := config.AppCfg().Env

	if appEnv == "production" {
		logger.Formatter = &logrus.JSONFormatter{}
	}

	logger.SetOutput(os.Stdout)

	if config.AppCfg().Debug {
		logger.SetLevel(logrus.DebugLevel)
	}
}

func GetLogger() *Logger {
	return logger
}
