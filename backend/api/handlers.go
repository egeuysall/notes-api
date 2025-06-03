package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

var notes = make(map[string]string)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "Welcome to Notes API v1. Available routes: POST /v1/notes, GET /v1/notes, GET /v1/notes/{id}, DELETE /v1/notes/{id}, GET /ping",
	}
	SendJson(w, response, http.StatusOK)
}

func checkPing(w http.ResponseWriter, r *http.Request) {
	SendJson(w, map[string]string{"status": "pong"}, http.StatusOK)
}

func createNote(w http.ResponseWriter, r *http.Request) {
	var note string
	err := json.NewDecoder(r.Body).Decode(&note)
	
	if err != nil {
		SendError(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	id := uuid.New().String()
	notes[id] = note

	response := map[string]interface{}{
		"notes": []Note{{Id: id, Note: note}},
	}
	SendJson(w, response, http.StatusCreated)
}

func getNote(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		SendError(w, "Missing note ID", http.StatusBadRequest)
		return
	}

	note, exists := notes[id]
	if !exists {
		SendError(w, "Note not found", http.StatusNotFound)
		return
	}

	SendJson(w, Note{Id: id, Note: note}, http.StatusOK)
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		SendError(w, "Missing note ID", http.StatusBadRequest)
		return
	}

	if _, exists := notes[id]; !exists {
		SendError(w, "Note not found", http.StatusNotFound)
		return
	}

	delete(notes, id)
	w.WriteHeader(http.StatusNoContent)
}