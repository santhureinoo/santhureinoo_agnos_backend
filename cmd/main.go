package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/santhureinoo_agnos_backend/configs"
	"github.com/santhureinoo_agnos_backend/middlewares"
	"github.com/santhureinoo_agnos_backend/routes"
)

func main() {

	// Load configuration
	configs.LoadConfig()

	// Create a new Gin router
	r := gin.Default()

	// Apply middleware
	r.Use(middlewares.LoggingMiddleware())

	// Set up routes
	routes.SetupRoutes(r)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}

}
