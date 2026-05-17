package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Email string `json:"email"`
}

type response struct {
	ID     uuid.UUID `json:"id"`
	CREATE time.Time `json:"created_at"`
	UPDATE time.Time `json:"updated_at"`
	EMAIL  string    `json:"email"`
}

func (a *ApiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
	userS := User{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userS)
	if err != nil {
		log.Println("Error occured with Decoding")
		respondWithError(w, 400, "Something went wrong")
		return
	}

	user, err := a.DB.CreateUser(r.Context(), userS.Email)

	payload := response{
		ID:     user.ID,
		CREATE: user.CreatedAt,
		UPDATE: user.UpdatedAt,
		EMAIL:  user.Email,
	}
	RespondWithJSON(w, 201, payload)
}

func (a *ApiConfig) HandleReset(w http.ResponseWriter, r *http.Request) {
	if a.PLATFORM == "dev" {
		respondWithError(w, 403, "Not allowed")
		return
	}
	err := a.DB.DeleteUsers(r.Context())

	if err != nil {
		log.Println("Error occured with Deleting all users")
		respondWithError(w, 400, "Something went wrong")
		return
	}

	type response struct {
		RESPONSE string `json:"response"`
	}

	payload := response{
		RESPONSE: "all users have been deleted",
	}
	RespondWithJSON(w, 200, payload)
}
