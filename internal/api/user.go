package api

import (
	"go-htmx/internal/locator"
	"go-htmx/internal/modules/auth"
	"go-htmx/internal/modules/user"

	"github.com/gin-gonic/gin"
)

func UserHandlers(r *gin.RouterGroup, locator locator.Locator) {
	service := locator.GetUserService()

	api := r.Group("/users").Use(auth.AuthMiddleware(locator.GetTokenManager()))
	uc := user.NewUserController(service)

	api.GET("/me", uc.GetMe)
	api.GET("/users", uc.GetUsers)
}
