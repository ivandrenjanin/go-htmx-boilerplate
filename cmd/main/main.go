package main

import (
	"go-htmx/internal/server"
	"go-htmx/pkg/config"
)

func main() {
	cfg := config.NewConfig()

	server.Init(cfg)
}
