// /core/middleware/jwt.go

package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4" // latest is v5, but ecosystem uses v4
)

var jwks *keyfunc.JWKS

// Loads the key system from keycloak
// and caches it for token auth.
func InitJWT() {
	jwksURL := os.Getenv("KEYCLOAK_JWKS")
	if jwksURL == "" {
		log.Fatal("KEYCLOAK_JWKS not set in .env")
	}

	options := keyfunc.Options{
		RefreshInterval: time.Hour,
		RefreshErrorHandler: func(err error) {
			log.Printf("JWKS refresh error: %v", err)
		},
	}

	var err error
	jwks, err = keyfunc.Get(jwksURL, options)
	if err != nil {
		log.Fatalf("Failed to load JWKS: %v", err)
	}
}

// AuthRequired uses Gin to check Authorization header for valid JWT.
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenStr, jwks.Keyfunc)
		if err != nil || !token.Valid {
			log.Printf("Invalid token: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT"})
			return
		}
		// optional: store claims in context
		c.Set("token", token)
		c.Next()
	}
}
