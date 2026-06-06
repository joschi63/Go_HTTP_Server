package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/joschi64/Go_HTTP_Server/internal/auth"
)

func (a *ApiConfig) HandleDeleteChirp(w http.ResponseWriter, r *http.Request) {
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

	id_string := r.PathValue("chirpID")

	id, err := uuid.Parse(id_string)

	if err != nil {
		respondWithError(w, 500, "Error converting given id into uuid")
		return
	}

	chirp, err := a.DB.GetChirp(r.Context(), uuid.UUID(id))

	if err != nil {
		respondWithError(w, 404, "Not found")
	}

	if userID != chirp.UserID.UUID {
		respondWithError(w, 403, "Forbidden")
		return
	}

	err = a.DB.DeleteChirp(r.Context(), id)

	if err != nil {
		respondWithError(w, 400, "something went wrong")
	}

	w.WriteHeader(204)
}
