package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// Logger is a middleware that logs the request method, path, status code, and latency
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Get request details
		method := c.Request.Method
		path := c.Request.URL.Path
		status := c.Writer.Status()

		// Log request details
		log.Printf("[%s] %s %s %d %s", method, path, c.ClientIP(), status, latency)
	}
}

// RequestID adds a unique request ID to each request
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate a unique request ID
		requestID := time.Now().UnixNano()
		
		// Set the request ID in the context
		c.Set("RequestID", requestID)
		
		// Add the request ID to the response headers
		c.Writer.Header().Set("X-Request-ID", string(requestID))
		
		c.Next()
	}
}