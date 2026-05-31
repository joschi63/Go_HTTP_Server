package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/joschi64/Go_HTTP_Server/internal/auth"
)

type LoginUser struct {
	Email      string `json:"email"`
	PASSWORD   string `json:"password"`
	EXPIRES_IN int64  `json:"expires_in_second"`
}

type LoginResponse struct {
	ID     uuid.UUID `json:"id"`
	CREATE time.Time `json:"created_at"`
	UPDATE time.Time `json:"updated_at"`
	EMAIL  string    `json:"email"`
	TOKEN  string    `json:"token"`
}

func (a *ApiConfig) HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	user := LoginUser{
		EXPIRES_IN: 3600,
	}

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

	token, err := auth.MakeJWT(db_user.ID, a.SECRET, time.Duration(user.EXPIRES_IN*int64(time.Second)))

	if err != nil {
		respondWithError(w, 500, "Error creating token")
		return
	}

	payload := LoginResponse{
		ID:     db_user.ID,
		CREATE: db_user.CreatedAt,
		UPDATE: db_user.UpdatedAt,
		EMAIL:  db_user.Email,
		TOKEN:  token,
	}
	RespondWithJSON(w, 200, payload)

}
