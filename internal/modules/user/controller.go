package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetMe(c *gin.Context)
	GetUsers(c *gin.Context)
}

type userController struct {
	userService UserService
}

func NewUserController(us UserService) UserController {
	return &userController{
		userService: us,
	}
}

func (uc *userController) GetMe(c *gin.Context) {
	id, ok := c.Get("userId")

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "User not found",
		})
		return
	}

	user, err := uc.userService.GetUser(id.(string))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (uc *userController) GetUsers(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid page parameter",
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid limit parameter",
		})
		return
	}

	users, err := uc.userService.GetUsersPaginated(page, limit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}
