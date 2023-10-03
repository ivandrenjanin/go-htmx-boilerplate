package db

import (
	"fmt"
	"go-htmx/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s sslmode=%s user=%s password=%s dbname=%s",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.SslMode,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
	)

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic(err.Error())
	}

	return db
}
