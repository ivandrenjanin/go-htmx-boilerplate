package main

import (
	"go-htmx/app/server"
	"go-htmx/pkg/config"
)

func main() {
	cfg := config.NewConfig()
	server.Init(cfg)
}
