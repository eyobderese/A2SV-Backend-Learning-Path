package infrastructure

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte("your_jwt_secret")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// JWT validation logic
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		claims, err := TotokenParser(authParts[1])

		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid JWT claims"})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		fmt.Println("clamis", claims)

		c.Next()
	}
}
