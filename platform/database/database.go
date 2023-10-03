package database

import (
	"fmt"

	"github.com/ivandrenjanin/go-fiber-htmx-boilerplate/pkg/config"
	_ "github.com/jackc/pgx/v5/stdlib" // load pgx driver for PostgreSQL
	"github.com/jmoiron/sqlx"
)

type DB struct{ *sqlx.DB }

var defaultDB = new(DB)

func (db *DB) connect(cfg *config.DB) error {
	dbURI := fmt.Sprintf("host=%s port=%d sslmode=%s user=%s password=%s dbname=%s",
		cfg.Host,
		cfg.Port,
		cfg.SslMode,
		cfg.User,
		cfg.Password,
		cfg.Name,
	)

	DB, err := sqlx.Connect("pgx", dbURI)

	db.DB = DB

	if err != nil {
		return err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConn)
	db.SetMaxIdleConns(cfg.MaxIdleConn)
	db.SetConnMaxLifetime(cfg.MaxConnLifetime)

	if err := db.Ping(); err != nil {
		defer db.Close()
		return fmt.Errorf("can not send ping to database, %w", err)
	}

	return nil
}

func GetDB() *DB {
	return defaultDB
}

func ConnectDB() error {
	return defaultDB.connect(config.DBCfg())
}
