package http

import (
	"encoding/json"
	"net/http"
)

type ResponseError struct {
	Error     string `json:"error"`
}

func SendResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	switch v := data.(type) {
	case error:
		data = ResponseError{v.Error()}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

