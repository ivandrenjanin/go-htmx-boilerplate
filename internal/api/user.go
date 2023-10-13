package api

import (
	"go-htmx/internal/locator"
	"go-htmx/internal/modules/auth"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserHandlers(r *gin.RouterGroup, locator locator.Locator) {
	tm := locator.GetTokenManager()
	service := locator.GetUserService()

	api := r.Group("/users").Use(auth.AuthMiddleware(tm))

	api.GET("/me", func(c *gin.Context) {
		id, ok := c.Get("userId")

		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "User not found",
			})
			return
		}

		user, err := service.GetUser(id.(string))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": user,
		})
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
}
