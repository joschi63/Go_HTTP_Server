package handler

import (
	"net/http"
	"sort"

	"github.com/google/uuid"
	"github.com/joschi64/Go_HTTP_Server/internal/database"
)

func (a *ApiConfig) HandleChirps(w http.ResponseWriter, r *http.Request) {

	author_id := r.URL.Query().Get("author_id")
	sortT := r.URL.Query().Get("sort")

	chirps := make([]database.Chirp, 0)

	if author_id != "" {
		userId, err := uuid.Parse(author_id)

		if err != nil {
			respondWithError(w, 400, "error parsing id")
		}

		chirps, err = a.DB.GetChirpsFromUser(r.Context(), uuid.NullUUID{
			UUID:  userId,
			Valid: true,
		})

		if err != nil {
			respondWithError(w, 500, "Error with getting all chirps")
			return
		}
	} else {
		var err error
		chirps, err = a.DB.GetChirps(r.Context())

		if err != nil {
			respondWithError(w, 500, "Error with getting all chirps")
			return
		}
	}

	if sortT == "desc" {
		sort.Slice(chirps, func(i, j int) bool { return chirps[i].CreatedAt.After(chirps[j].CreatedAt) })
	}

	payload := make([]responseChirp, 0, len(chirps))

	for _, chirp := range chirps {
		result := responseChirp{
			ID:     chirp.ID,
			CREATE: chirp.CreatedAt,
			UPDATE: chirp.UpdatedAt,
			BODY:   chirp.Body,
			USERID: chirp.UserID,
		}
		payload = append(payload, result)
	}

	respondWithJSON(w, 200, payload)
}
