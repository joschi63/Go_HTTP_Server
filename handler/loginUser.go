package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/joschi64/Go_HTTP_Server/internal/auth"
	"github.com/joschi64/Go_HTTP_Server/internal/database"
)

type LoginUser struct {
	Email    string `json:"email"`
	PASSWORD string `json:"password"`
}

type LoginResponse struct {
	ID          uuid.UUID `json:"id"`
	CREATE      time.Time `json:"created_at"`
	UPDATE      time.Time `json:"updated_at"`
	EMAIL       string    `json:"email"`
	IsChripyRed bool      `json:"is_chirpy_red"`
	TOKEN       string    `json:"token"`
	REFRESH     string    `json:"refresh_token"`
}

func (a *ApiConfig) HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	user := LoginUser{}

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

	token, err := auth.MakeJWT(db_user.ID, a.SECRET, time.Duration(time.Hour*1))

	if err != nil {
		respondWithError(w, 500, "Error creating token")
		return
	}

	refresh_token := auth.MakeRefreshToken()

	a.DB.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
		Token: refresh_token,
		UserID: uuid.NullUUID{
			UUID:  db_user.ID,
			Valid: true,
		},
		ExpiresAt: sql.NullTime{
			Time:  time.Now().Add(60 * 24 * time.Hour),
			Valid: true,
		},
	})

	payload := LoginResponse{
		ID:          db_user.ID,
		CREATE:      db_user.CreatedAt,
		UPDATE:      db_user.UpdatedAt,
		EMAIL:       db_user.Email,
		IsChripyRed: db_user.IsChirpyRed.Bool,
		TOKEN:       token,
		REFRESH:     refresh_token,
	}
	respondWithJSON(w, 200, payload)

}
