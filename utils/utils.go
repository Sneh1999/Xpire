package utils

import (
	"encoding/json"
	"net/http"
)

//WritePretty helps in setting up the response
func WritePretty(w http.ResponseWriter, statusCode int, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		// TODO: check for internal server error
		w.WriteHeader(http.StatusInternalServerError)
	}
}
