package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dp487/scaling-waddle/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get server host and port from environment variables
	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")
	if host == "" || port == "" {
		log.Fatal("SERVER_HOST or SERVER_PORT not set in environment")
	}

	// Create a new Gorilla Mux router
	r := mux.NewRouter()

	// Register the routes for the bookstore
	routes.RegisterBookStoreRoutes(r)

	// Start the server
	serverAddress := host + ":" + port
	log.Printf("Server starting on %s", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, r))
}
