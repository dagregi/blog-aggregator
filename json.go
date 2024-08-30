package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error: %v couldn't marshal JSON: %v\n", err, payload)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(statusCode)
	w.Write(payloadJSON)
}

func respondWithError(w http.ResponseWriter, statusCode int, msg string) {
	if statusCode > 499 {
		log.Printf("Resopnding with 5XX error: %s", msg)
	}
	respondWithJSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: msg,
	})
}
