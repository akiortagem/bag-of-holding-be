package middelwares

import (
	"net/http"
	"strings"

	"github.com/akiortagem/bag-of-holding-be/internal/core/auth/usecase"
	"github.com/akiortagem/bag-of-holding-be/internal/core/config"
	"github.com/gin-gonic/gin"
)

func AuthRequired(cfg config.ApiConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "missing or invalid authorization header",
			})
			c.Abort()
			return
		}

		token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "missing bearer token",
			})
			c.Abort()
			return
		}

		var userID int64

		userID, err := usecase.ValidateJWT(token, cfg.Secret, cfg.JWTIssuer)

		if err != nil {
			c.Status(401)
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
