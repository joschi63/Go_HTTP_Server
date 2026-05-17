package handler

import "net/http"

func (a *ApiConfig) HandleChirps(w http.ResponseWriter, r *http.Request) {

	chirps, err := a.DB.GetChirps(r.Context())

	if err != nil {
		respondWithError(w, 500, "Error with getting all chirps")
		return
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

	RespondWithJSON(w, 200, payload)
}
