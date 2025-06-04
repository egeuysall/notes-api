package main

import (
	"log"
	"net/http"

	"github.com/egeuysall/notes-api/api"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Define the router
	router := api.NewRouter()

	// Start the server
	log.Printf("Server starting on http://localhost:8080")
	err = http.ListenAndServe(":8080", router)

	if err != nil {
		log.Fatal(err)
	}
}