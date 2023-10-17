package auth

import (
	"go-htmx/internal/modules/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController interface {
	SignIn(c *gin.Context)
	SignUp(c *gin.Context)
}

type authController struct {
	userService user.UserService
	tm          TokenManager
}

func NewAuthController(us user.UserService, tm TokenManager) AuthController {
	return &authController{
		userService: us,
		tm:          tm,
	}
}

func (ac *authController) SignIn(c *gin.Context) {
	var dto LoginRequestDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := ac.userService.GetUserByEmail(dto.Email)

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

	t, err := ac.tm.GenerateToken(u.ID)

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
}

func (ac *authController) SignUp(c *gin.Context) {
	var dto user.CreateUserRequestDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ac.userService.CreateUser(&dto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
