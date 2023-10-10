package route

import (
	"go-htmx/internal/locator"
	"go-htmx/internal/modules/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ApiHandlers(r *gin.Engine, locator locator.Locator) {
	api := r.Group("/api/v1")
	service := locator.GetUserService()
	sessionService := locator.GetSessionService()

	api.POST("/users", func(c *gin.Context) {
		var dto user.CreateUserRequestDTO
		c.Bind(&dto)

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

	api.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")

		user, err := service.GetUser(id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, user)
	})

	api.GET("/users", func(c *gin.Context) {
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

		users, err := service.GetUsersPaginated(page, limit)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, users)
	})

	type LoginRequestDTO struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8,max=32"`
	}

	api.POST("/login", func(c *gin.Context) {
		var dto LoginRequestDTO
		c.Bind(&dto)

		u, err := service.GetUserByEmail(dto.Email)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		if u.Password != dto.Password {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid credentials",
			})
			return
		}

		sessionService.SetKey(c, u.ID, u)

		c.JSON(http.StatusOK, gin.H{
			"data": u,
		})
	})
}
