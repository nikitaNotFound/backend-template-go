package utils

import (
	"encoding/json"
	"net/http"
)

type ApiError struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
}

func WriteClientError(w http.ResponseWriter, err ApiError) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	return json.NewEncoder(w).Encode(err)
}
