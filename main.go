package main

import (
	"fmt"
	"log"

	"github.com/fitnis/patient-service/configs"
	"github.com/fitnis/patient-service/internal/database"
	"github.com/fitnis/patient-service/internal/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load configuration
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Connect to database
	if err := database.ConnectDB(config); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create a new Echo instance
	e := echo.New()

	// Setup routes
	routes.SetupRoutes(e, config)

	// Start server
	log.Printf("Starting server on port %d", config.Port)
	if err := e.Start(fmt.Sprintf(":%d", config.Port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
