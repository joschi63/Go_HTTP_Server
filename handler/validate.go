package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func HandleValidateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Println("Error occured with Decoding")
		respondWithError(w, 400, "Something went wrong")
		return
	}

	if len(params.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long")
		return
	}
	words := strings.Split(params.Body, " ")

	for i, word := range words {
		if strings.ToLower(word) == "kerfuffle" || strings.ToLower(word) == "sharbert" ||
			strings.ToLower(word) == "fornax" {
			words[i] = "****"
		}
	}

	type positiveResponse struct {
		Cleaned_Body string `json:"cleaned_body"`
	}

	payload := positiveResponse{
		Cleaned_Body: strings.Join(words, " "),
	}
	respondWithJSON(w, 200, payload)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {

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

	respondWithJSON(w, code, body)
}
