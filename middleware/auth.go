package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()

		if !ok || username != "admin" || password != "mase1234" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Username dan atau Password salah"})
			c.Abort()
			return
		}

		c.Next()
	}
}
