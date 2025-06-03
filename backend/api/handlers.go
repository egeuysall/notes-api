package api

import (
	"encoding/json"
	"net/http"
)

// Handle the root route to welcome users
func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"message": "Welcome to Notes API v1. Available routes: POST /v1/note, GET /v1/note/{id}, DELETE /v1/note/{id}",
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		errorResponse := map[string]string{
			"error": err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
}

// Handle the ping route to check if the server is alive
func checkPing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"status": "pong",
	}

	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		errorResponse := map[string]string{
			"error": err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
}

func createNote(w http.ResponseWriter, r *http.Request) {

}

func getNote(w http.ResponseWriter, r *http.Request) {

}

func deleteNote(w http.ResponseWriter, r *http.Request) {

}
