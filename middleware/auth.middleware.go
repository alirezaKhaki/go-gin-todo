package middleware

import (
	"net/http"
	"strings"

	models "github.com/alirezaKhaki/go-gin/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the JWT token from the request header
		tokenString := c.GetHeader("Authorization")

		// Check if the token is missing or empty
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or empty token"})
			c.Abort()
			return
		}

		// Check for the "Bearer " prefix and remove it
		parts := strings.SplitN(tokenString, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		tokenString = parts[1]

		// Parse and validate the token
		token, err := jwt.ParseWithClaims(tokenString, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("your-secret-key"), nil // Use your JWT secret here
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Check if the token is valid
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is not valid"})
			c.Abort()
			return
		}

		// Extract the claims from the token
		claims, ok := token.Claims.(*models.UserClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to extract claims"})
			c.Abort()
			return
		}

		// You can now access the authenticated user's information through 'claims'
		// For example, you can access claims.UserID to identify the user

		// Attach the claims to the context for use in downstream handlers
		c.Set("user", claims)

		// Continue processing the request
		c.Next()
	}
}
