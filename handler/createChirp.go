package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/joschi64/Go_HTTP_Server/internal/database"
)

type responseChirp struct {
	ID     uuid.UUID     `json:"id"`
	CREATE time.Time     `json:"created_at"`
	UPDATE time.Time     `json:"updated_at"`
	BODY   string        `json:"body"`
	USERID uuid.NullUUID `json:"user_id"`
}

func (a *ApiConfig) HandleChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body   string    `json:"body"`
		UserID uuid.UUID `json:"user_id"`
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
	cleaned_body := strings.Join(words, " ")

	chirp, err := a.DB.CreateChirp(r.Context(), database.CreateChirpParams{
		Body: cleaned_body,
		UserID: uuid.NullUUID{
			UUID:  params.UserID,
			Valid: true,
		},
	})
	if err != nil {
		respondWithError(w, 500, "Something went wrong")
		return
	}

	payload := responseChirp{
		ID:     chirp.ID,
		CREATE: chirp.CreatedAt,
		UPDATE: chirp.UpdatedAt,
		BODY:   chirp.Body,
		USERID: chirp.UserID,
	}
	RespondWithJSON(w, 201, payload)
}
