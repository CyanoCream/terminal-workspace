package main

import (
	"log"
	"terminal/api/internal/app"
)

func main() {
	// Initialize the application
	application, err := app.NewApp()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// Start the server
	log.Printf("Starting server on port %s", application.Server.Config().Network)
	if err := application.Server.Listen(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
