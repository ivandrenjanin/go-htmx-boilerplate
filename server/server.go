package server

import (
	"go-htmx/config"
	"go-htmx/db"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	Echo   *echo.Echo
	DB     *gorm.DB
	Config *config.Config
}

func newServer(cfg *config.Config) *Server {
	return &Server{
		Echo:   echo.New(),
		DB:     db.Init(cfg),
		Config: cfg,
	}
}

func (s *Server) Start(port string) error {
	return s.Echo.Start(port)
}

func Init(cfg *config.Config) {
	app := newServer(cfg)

	err := app.Start(cfg.HTTP.Port)

	if err != nil {
		log.Fatalf("Port already used %v", err)
	}
}
