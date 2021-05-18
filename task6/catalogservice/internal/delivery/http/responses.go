package http

import (
	"encoding/json"
	"net/http"
)

type ResponseError struct {
	Error     string `json:"error"`
}

type PaginatedResults struct {
	Next      string      `json:"next"      example:"?page=5"`
	Previous  string      `json:"previous"  example:"?page=3"`
	Page      uint64      `json:"page"      example:"4"`
	PageSize  uint64      `json:"page_size" example:"10"`
	Results   interface{} `json:"results"`
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
