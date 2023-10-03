package main

import (
	"github.com/ivandrenjanin/go-fiber-htmx-boilerplate/cmd/server"
	"github.com/ivandrenjanin/go-fiber-htmx-boilerplate/pkg/config"
)

func main() {

	// setup various configuration for app
	config.LoadAllConfigs(".env")

	server.Serve()
}
