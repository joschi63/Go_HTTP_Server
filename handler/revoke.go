package handler

import (
	"net/http"

	"github.com/joschi64/Go_HTTP_Server/internal/auth"
)

func (a *ApiConfig) HandleRevoke(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)

	if err != nil {
		respondWithError(w, 401, "Error getting refresh token from header")
		return
	}

	err = a.DB.UpdateRevokeRefreshToken(r.Context(), token)

	if err != nil {
		respondWithError(w, 401, "Error revoking refresh token")
		return
	}

	w.WriteHeader(204)
}
