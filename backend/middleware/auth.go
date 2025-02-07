package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implement JWT validation here
		// For now, just set a dummy user ID
		c.Set("userID", uint(1))
		c.Next()
	}
}
