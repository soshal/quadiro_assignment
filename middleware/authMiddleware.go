package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

// AuthMiddleware is a middleware for authentication
func AuthMiddleware(c *gin.Context) {
    token := c.GetHeader("Authorization")

    if token == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
        c.Abort()
        return
    }

    if !strings.HasPrefix(token, "Bearer ") {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
        c.Abort()
        return
    }

    token = strings.TrimPrefix(token, "Bearer ")

    // This is where you would normally validate the token
    // For this example, we will just check if the token is "admin_token"
    if token != "admin_token" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        c.Abort()
        return
    }

    // If token is valid, proceed to the next middleware/handler
    c.Next()
}
