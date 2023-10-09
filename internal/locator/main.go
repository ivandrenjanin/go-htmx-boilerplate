package locator

import (
	"database/sql"
	"go-htmx/internal/modules/user"
	"go-htmx/pkg/config"
)

func NewLocator(db *sql.DB, cfg *config.Config) Locator {
	return &locator{db, cfg}
}

type Locator interface {
	GetUserService() user.UserService
	GetConfig() *config.Config
	GetDB() *sql.DB
}

type locator struct {
	db  *sql.DB
	cfg *config.Config
}

func (l *locator) GetUserService() user.UserService {
	return user.NewUserService(l.db)
}

func (l *locator) GetConfig() *config.Config {
	return l.cfg
}

func (l *locator) GetDB() *sql.DB {
	return l.db
}
