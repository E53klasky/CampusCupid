package middleware

import (
	"os"      // For accessing environment variables
	"strings" // For string manipulation (trimming prefix)

	"github.com/gin-gonic/gin"     // Gin web framework
	"github.com/golang-jwt/jwt/v5" // JWT handling library
)

// AuthMiddleware is a middleware function that handles JWT-based authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header from the incoming request
		authHeader := c.GetHeader("Authorization")
		// If the Authorization header is missing, return a 401 Unauthorized response
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort() // Stop further processing of the request
			return
		}

		// Extract the token from the Authorization header (expected format: "Bearer <token>")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		// If the token format is incorrect (e.g., missing "Bearer "), return an error
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "Invalid token format"})
			c.Abort() // Stop further processing of the request
			return
		}

		// Parse the token and validate it using the JWT secret key from environment variables
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Return the secret key used for signing the JWT (from environment variable)
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		// If there is an error parsing the token or it's invalid, return a 401 Unauthorized error
		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort() // Stop further processing of the request
			return
		}

		// Extract the claims (payload) from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		// If extracting claims fails or the claims are invalid, return an error
		if !ok {
			c.JSON(401, gin.H{"error": "Invalid token claims"})
			c.Abort() // Stop further processing of the request
			return
		}

		// Extract the user ID from the claims (JWT numbers are typically stored as float64)
		userID, ok := claims["user_id"].(float64)
		// If the user ID is not found or is in an incorrect format, return an error
		if !ok {
			c.JSON(401, gin.H{"error": "Invalid user ID in token"})
			c.Abort() // Stop further processing of the request
			return
		}

		// Set the userID in the context to make it available for other handlers in the request chain
		c.Set("userID", uint(userID))

		// Continue to the next handler in the chain
		c.Next()
	}
}
