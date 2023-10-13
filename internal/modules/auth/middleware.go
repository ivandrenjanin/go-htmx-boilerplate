package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(tm TokenManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		userId, err := tm.ValidateToken(token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		c.Set("userId", userId)
		c.Next()
	}
}
