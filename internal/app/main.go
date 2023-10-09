package app

import (
	"go-htmx/internal/locator"
	"go-htmx/internal/server"
	"go-htmx/pkg/config"
	"go-htmx/pkg/db"
	"log"
)

func Init() {
	cfg := config.NewConfig()
	db, err := db.Init(cfg)

	if err != nil {
		log.Fatalf("db init error: %s\n", err)
		return
	}

	_, err = db.Exec(`
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

	CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(), 
		name VARCHAR(255) NOT NULL, 
		email VARCHAR(255) NOT NULL UNIQUE, 
		password VARCHAR(255) NOT NULL, 
		created_at TIMESTAMP DEFAULT NOW(), 
		updated_at TIMESTAMP DEFAULT NOW(), 
		deleted_at TIMESTAMP DEFAULT NULL
	)`)

	if err != nil {
		log.Fatalf("db init error: %s\n", err)
		return
	}

	locator := locator.NewLocator(db, cfg)

	server.Init(locator)
}
