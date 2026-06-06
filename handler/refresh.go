package handler

import (
	"net/http"
	"time"

	"github.com/joschi64/Go_HTTP_Server/internal/auth"
)

type RefreshResponse struct {
	TOKEN string `json:"token"`
}

func (a *ApiConfig) HandleRefresh(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)

	if err != nil {
		respondWithError(w, 400, "Error getting refresh token from header")
		return
	}

	db_token, err := a.DB.GetRefreshToken(r.Context(), token)

	if err != nil {
		respondWithError(w, 400, "Error getting refresh token from db")
		return
	}

	if time.Now().After(db_token.ExpiresAt.Time) || db_token.RevokedAt.Valid == true {
		respondWithError(w, 401, "Error refresh token out of date ")
		return
	}

	jwt_token, err := auth.MakeJWT(db_token.UserID.UUID, a.SECRET, time.Duration(time.Hour*1))

	if err != nil {
		respondWithError(w, 400, "Error creating access token")
		return
	}

	payload := RefreshResponse{
		TOKEN: jwt_token,
	}

	respondWithJSON(w, 200, payload)
}
