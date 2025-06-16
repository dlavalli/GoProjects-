package main

import (
	"JunieSampleProject/internal/db"
	"JunieSampleProject/internal/handlers"
	"JunieSampleProject/internal/middleware"
	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	//"time"
)

func main() {
	// Create a new gin router
	router := gin.New()

	// Add recovery middleware to recover from panics
	router.Use(gin.Recovery())

	// Add custom middleware
	router.Use(middleware.Logger())
	router.Use(middleware.RequestID())

	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	// Serve static files from the web directory
	router.StaticFile("/", "./web/index.html")
	router.Static("/static", "./web")

	// Initialize the user store
	userStore := db.NewUserStore()

	// Initialize the user handler
	userHandler := handlers.NewUserHandler(userStore)

	// Define a simple health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Service is up and running",
		})
	})

	// API v1 routes group
	v1 := router.Group("/api/v1")
	{
		// Example endpoints
		v1.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello from the API!",
			})
		})

		// Register user routes
		userHandler.RegisterRoutes(v1)
	}

	// Start the server
	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
