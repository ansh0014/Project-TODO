package main

import (
	"log"
	"net/http"

	"Todo/config"
	"Todo/routes"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Get the port from environment or fallback to default
	port := config.GetEnv("PORT", "8080")

	// Register application routes
	router := routes.RegisterRoutes()

	// Start the HTTP server
	log.Println("🚀 Server running on port:", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
