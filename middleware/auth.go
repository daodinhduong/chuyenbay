package middleware

import (
	"go-api/internal/token"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		if tokenString == "" {
			c.JSON(401, gin.H{
				"error": "request doesn't contain access token",
			})
			c.Abort()
			return
		}

		if err := token.ValidateToken(tokenString); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
