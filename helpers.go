package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type errResponse struct {
	Error string `json:"error"`
}

func setErrorResponse(w http.ResponseWriter, code int, msg string) {
	// validate only internal errors
	if code > 499 {
		log.Println("Responding with 5XX error: ", msg)
	}
	setJSONResponse(w, code, errResponse{
		Error: msg,
	})
}

func setJSONResponse(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}