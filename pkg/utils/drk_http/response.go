package drk_http

import (
	"encoding/json"
	"net/http"
)

func NewResponseWithStatus(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}

func NewError(w http.ResponseWriter, status int, err error) {
	NewResponseWithStatus(w, status, map[string]string{"error": err.Error()})
}
