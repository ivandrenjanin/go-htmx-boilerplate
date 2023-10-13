package api

import (
	"go-htmx/internal/locator"
	"go-htmx/internal/modules/auth"
	"go-htmx/internal/modules/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func AuthHandlers(r *gin.RouterGroup, locator locator.Locator) {
	api := r.Group("/auth")
	service := locator.GetUserService()
	tm := locator.GetTokenManager()

	api.POST("/sign-up", func(c *gin.Context) {
		var dto user.CreateUserRequestDTO

		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := service.CreateUser(&dto)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})

	api.POST("/sign-in", func(c *gin.Context) {
		var dto auth.LoginRequestDTO

		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		u, err := service.GetUserByEmail(dto.Email)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(dto.Password))

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid credentials",
			})
			return
		}

		t, err := tm.GenerateToken(u.ID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":  u,
			"token": t,
		})
	})

}
