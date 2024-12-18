package utils

import (
	"encoding/json"
	"net/http"
)

func WriteOK(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(data)
}
