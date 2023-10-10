package locator

import (
	"database/sql"
	"go-htmx/internal/modules/user"
	"go-htmx/pkg/config"
	"go-htmx/pkg/session"
	"sync"
)

type Locator interface {
	GetConfig() *config.Config
	GetDB() *sql.DB
	GetSessionService() session.Session
	GetUserService() user.UserService
}

type locator struct {
	db             *sql.DB
	cfg            *config.Config
	sessionService session.Session
	userService    user.UserService
}

var lock = &sync.Mutex{}
var instance *locator

func NewLocator(db *sql.DB, cfg *config.Config) Locator {
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = &locator{
			db:             db,
			cfg:            cfg,
			sessionService: session.NewSession(),
			userService:    user.NewUserService(db),
		}
	}

	return instance
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

func (l *locator) GetSessionService() session.Session {
	return l.sessionService
}
