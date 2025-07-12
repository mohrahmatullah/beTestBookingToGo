package response

import (
	"encoding/json"
	"net/http"
)

// Success membuat response JSON success
func Success(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "success",
		"data":    data,
	})
}

// Error membuat response JSON error
func Error(w http.ResponseWriter, statusCode int, errMessage string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "error",
		"data":    nil,
		"error":   errMessage,
	})
}