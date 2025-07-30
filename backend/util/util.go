package util

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Error writing payload: %v", err)
	}
}

func RespondWithError(w http.ResponseWriter, code int, message string, err error) {
	type errorPayload struct {
		Error string `json:"error"`
	}

	resp := errorPayload{
		Error: fmt.Sprintf("%s: %v", message, err),
	}

	data, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Error marshalling error payload: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Error writing error payload: %v", err)
	}
}
