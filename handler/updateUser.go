package handler

import (
	"encoding/json"
	"net/http"

	"github.com/joschi64/Go_HTTP_Server/internal/auth"
	"github.com/joschi64/Go_HTTP_Server/internal/database"
)

type updateInput struct {
	EMAIL    string `json:"email"`
	PASSWORD string `json:"password"`
}

func (a *ApiConfig) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)

	if err != nil {
		respondWithError(w, 401, "Something went wrong")
		return
	}

	userID, err := auth.ValidateJWT(token, a.SECRET)

	if err != nil {
		respondWithError(w, 401, "Something went wrong")
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := updateInput{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 398, "Something went wrong")
		return
	}

	hashedInputPassword, err := auth.HashPassword(params.PASSWORD)

	if err != nil {
		respondWithError(w, 401, "Unauthorized")
		return
	}

	err = a.DB.UpdateUser(r.Context(), database.UpdateUserParams{
		Email:          params.EMAIL,
		HashedPassword: hashedInputPassword,
		ID:             userID,
	})

	if err != nil {
		respondWithError(w, 400, "Something went wrong")
		return
	}

	user, err := a.DB.GetUser(r.Context(), params.EMAIL)

	if err != nil {
		respondWithError(w, 400, "Something went wrong")
		return
	}

	payload := response{
		EMAIL:  user.Email,
		CREATE: user.CreatedAt,
		UPDATE: user.UpdatedAt,
		ID:     user.ID,
	}

	respondWithJSON(w, 200, payload)

}
