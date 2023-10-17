package api

import (
	"go-htmx/internal/locator"
	"go-htmx/internal/modules/auth"

	"github.com/gin-gonic/gin"
)

func AuthHandlers(r *gin.RouterGroup, locator locator.Locator) {
	api := r.Group("/auth")
	ac := auth.NewAuthController(locator.GetUserService(), locator.GetTokenManager())

	api.POST("/sign-up", ac.SignUp)
	api.POST("/sign-in", ac.SignIn)

}
