package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(
	w http.ResponseWriter,
	code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)

}

func respondWithError(
	w http.ResponseWriter,
	code int,
	msg string) {
	if code > 499 {
		log.Println("RespondWithError:", msg)

	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}
