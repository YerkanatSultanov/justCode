package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var authorizedToken = "secretKey"

func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != authorizedToken {
		c.JSON(http.StatusForbidden, gin.H{"error": "403 Unauthorized"})
		c.Abort()
		return
	}
	c.Next()
}
