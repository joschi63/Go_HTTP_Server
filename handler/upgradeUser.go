package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/joschi64/Go_HTTP_Server/internal/auth"
)

type request struct {
	EVENT string `json:"event"`
	DATA  struct {
		USERID string `json:"user_id"`
	} `json:"data"`
}

func (a *ApiConfig) HandleUpgradeUser(w http.ResponseWriter, r *http.Request) {
	req := request{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		log.Println("Error occured with Decoding")
		respondWithError(w, 400, "Something went wrong")
		return
	}

	if req.EVENT != "user.upgraded" {
		w.WriteHeader(204)
		return
	}

	apiKey, err := auth.GetAPIKey(r.Header)

	if err != nil || apiKey != a.PolkaKey {
		respondWithError(w, 401, "wrong APIKey")
		return
	}

	userID, err := uuid.Parse(req.DATA.USERID)

	if err != nil {
		respondWithError(w, 400, "something went wrong")
		return
	}

	err = a.DB.UpgradeUser(r.Context(), userID)

	if err != nil {
		respondWithError(w, 404, "something went wrong")
		return
	}

	w.WriteHeader(204)
}
