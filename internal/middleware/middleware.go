package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/glng-swndru/simple-forum/internal/configs"
	"github.com/glng-swndru/simple-forum/pkg/jwt"
)

// Middleware untuk autentikasi JWT
func AuthMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJWT
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")

		header = strings.TrimSpace(header)
		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		userID, username, err := jwt.ValidateToken(header, secretKey)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		c.Set("userID", userID)
		c.Set("username", username)
		c.Next()
	}
}

func AuthRefreshMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJWT
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")

		header = strings.TrimSpace(header)
		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		userID, username, err := jwt.ValidateTokenWithoutExpiry(header, secretKey)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		c.Set("userID", userID)
		c.Set("username", username)
		c.Next()
	}
}
