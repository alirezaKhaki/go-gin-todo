package middleware

import (
	"net/http"
	"strings"

	"github.com/alirezaKhaki/go-gin/domain"
	"github.com/alirezaKhaki/go-gin/lib"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware middleware for jwt authentication
type JWTAuthMiddleware struct {
	service domain.AuthService
	logger  lib.Logger
}

// NewJWTAuthMiddleware creates new jwt auth middleware
func NewJWTAuthMiddleware(
	logger lib.Logger,
	service domain.AuthService,
) JWTAuthMiddleware {
	return JWTAuthMiddleware{
		service: service,
		logger:  logger,
	}
}

// Setup sets up jwt auth middleware
func (m JWTAuthMiddleware) Setup() {}

// Handler handles middleware functionality
func (m JWTAuthMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := m.service.Authorize(authToken)
			if authorized.Valid {

				claims, ok := authorized.Claims.(jwt.MapClaims)
				if !ok {
					c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token claims"})
					c.Abort()
					return
				}

				// Get user information from the token
				userID, ok := claims["id"].(float64)
				if !ok {
					c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid user ID"})
					c.Abort()
					return
				}

				// Attach user information to the context for use in handlers
				c.Set("id", int(userID))
				c.Next()

				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			m.logger.Error(err)
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "you are not authorized",
		})
		c.Abort()
	}
}
