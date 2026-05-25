package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/joschi64/Go_HTTP_Server/internal/auth"
)

func (a *ApiConfig) HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	user := User{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		log.Println("Error occured with Decoding")
		respondWithError(w, 400, "Something went wrong")
		return
	}

	db_user, err := a.DB.GetUser(r.Context(), user.Email)

	if err != nil {
		respondWithError(w, 401, "Incorrect email or password (db)")
		return
	}

	correct_login, err := auth.CheckPasswordHash(user.PASSWORD, db_user.HashedPassword)

	if err != nil {
		respondWithError(w, 400, "Something went wrong while checking password")
		return
	}

	if !correct_login {
		respondWithError(w, 401, "Incorrect email or password")
		return
	}

	payload := response{
		ID:     db_user.ID,
		CREATE: db_user.CreatedAt,
		UPDATE: db_user.UpdatedAt,
		EMAIL:  db_user.Email,
	}
	RespondWithJSON(w, 200, payload)

}
