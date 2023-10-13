package locator

import (
	"database/sql"
	"go-htmx/internal/modules/auth"
	"go-htmx/internal/modules/user"
	"go-htmx/pkg/config"
)

type Locator interface {
	GetConfig() *config.Config
	GetDB() *sql.DB
	GetUserService() user.UserService
	GetTokenManager() auth.TokenManager
}

type locator struct {
	db           *sql.DB
	cfg          *config.Config
	tokenManager auth.TokenManager
	userService  user.UserService
}

func NewLocator(db *sql.DB, cfg *config.Config) Locator {
	return &locator{
		db:           db,
		cfg:          cfg,
		tokenManager: auth.NewTokenManager(cfg.App.AccessSecret, 10),
		userService:  user.NewUserService(db),
	}
}

func (l *locator) GetUserService() user.UserService {
	return l.userService
}

func (l *locator) GetConfig() *config.Config {
	return l.cfg
}

func (l *locator) GetDB() *sql.DB {
	return l.db
}

func (l *locator) GetTokenManager() auth.TokenManager {
	return l.tokenManager
}
