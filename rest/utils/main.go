package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type ApiError struct {
	Error string `json:"error"`
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

/*
Util function to wrap HTTP handler functions for central error handling
*/
func MakeHTTPHandlerFunc(apiFunc apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFunc(w, r); err != nil {

			log.Printf("ERROR: %s", err.Error())

			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})

		}
	}
}

/*
Util function to write json response to HTTP response writer
*/
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
