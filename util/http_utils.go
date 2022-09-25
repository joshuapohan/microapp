package util

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

func RespondWithError(w http.ResponseWriter, status int, err Error) {

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

func ResponseJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
