package session

import (
	"database/sql"
	"go-htmx/pkg/config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/postgres"
	"github.com/gin-gonic/gin"
)

type Session interface {
	GetSession(c *gin.Context) sessions.Session
	GetKey(c *gin.Context, key string) interface{}
	SetKey(c *gin.Context, key string, data interface{})
}

type session struct{}

func NewSession() Session {
	return &session{}
}

func (s *session) GetSession(c *gin.Context) sessions.Session {
	return sessions.Default(c)
}

func (s *session) GetKey(c *gin.Context, key string) interface{} {
	session := s.GetSession(c)
	return session.Get(key)
}

func (s *session) SetKey(c *gin.Context, key string, data interface{}) {
	session := s.GetSession(c)
	session.Set(key, data)
	session.Save()
}

func NewSessionMiddleware(db *sql.DB, cfg *config.Config) (gin.HandlerFunc, error) {
	store, err := postgres.NewStore(db, []byte(cfg.App.AccessSecret))

	if err != nil {
		return nil, err
	}

	store.Options(sessions.Options{
		MaxAge: 86400 * 7,
		// Secure:   true,
		HttpOnly: true,
	})

	return sessions.Sessions("pgsession", store), nil
}
