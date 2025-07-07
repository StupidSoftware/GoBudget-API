package middleware

import (
	"net/http"
	"strings"

	"github.com/breno5g/GoBudget/config"
	"github.com/breno5g/GoBudget/internal/utils"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		logger := config.GetLogger("auth")
		logger.Infof("Auth header: %s", authHeader)

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
			c.Abort()
			return
		}

		token, err := utils.DecodeToken(strings.TrimPrefix(authHeader, "Bearer "))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("user_id", token)
		c.Next()
	}
}
