package main

import (
	"go-htmx/config"
	"go-htmx/server"
)

func main() {
	cfg := config.NewConfig()
	server.Init(cfg)
}
