package api

import (
	"encoding/json"
	"net/http"
)

type Note struct {
	Id   string `json:"id"`
	Note string `json:"note"`
}

func SendJson(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SendError(w http.ResponseWriter, message string, statusCode int) {
	SendJson(w, map[string]string{"error": message}, statusCode)
}