package response

import (
	"encoding/json"
	"net/http"
)


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