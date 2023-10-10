package db

import (
	"database/sql"
	"fmt"
	"go-htmx/pkg/config"

	_ "github.com/lib/pq"
)

func Init(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.Port,
		cfg.Database.Sslmode,
	)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
