package handler

import (
	"net/http"

	"github.com/google/uuid"
)

func (a *ApiConfig) GetChirp(w http.ResponseWriter, r *http.Request) {

	id_string := r.PathValue("chirpID")

	id, err := uuid.Parse(id_string)

	if err != nil {
		respondWithError(w, 500, "Error converting given id into uuid")
		return
	}

	chirp, err := a.DB.GetChirp(r.Context(), uuid.UUID(id))

	if err != nil {
		respondWithError(w, 404, "Searched chirp not found")
		return
	}

	payload := responseChirp{
		ID:     chirp.ID,
		CREATE: chirp.CreatedAt,
		UPDATE: chirp.UpdatedAt,
		BODY:   chirp.Body,
		USERID: chirp.UserID,
	}

	RespondWithJSON(w, 200, payload)
	return
}
