package util

import (
	"encoding/json"
	"net/http"
)

// ResponseOk sends json formatted content to http response
func ResponseOk(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(body)
}

// ResponseError sends json formatted error to http response
func ResponseError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "applicaiton/json")

	body := map[string]string{
		"error": message,
	}
	json.NewEncoder(w).Encode(body)
}
