package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {

	body := payload

	data, err := json.Marshal(body)

	if err != nil {
		log.Printf("Error mashalling JSON-Error: %s", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)

}
func respondWithError(w http.ResponseWriter, code int, msg string) {
	type returnVals struct {
		Error string `json:"error"`
	}
	body := returnVals{
		Error: msg,
	}

	RespondWithJSON(w, code, body)
}
