package main

import (
	"log"
	"net/http"
	"os"

	"github.com/user/blackroot/internal/api"
	"github.com/user/blackroot/internal/logger"
)

func main() {
	// Setup basic logging - nothing too fancy for hackathon
	logger.InitLogger()
	log.Println("Starting Blackroot AI-Powered Sandbox Platform...")

	// init API routes
	router := api.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default for local testing
	}

	log.Printf("Server listening on port %s", port)
	
	// TODO: add graceful shutdown if time permits
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
