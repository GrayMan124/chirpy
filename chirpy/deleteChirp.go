package main

import (
	"github.com/GrayMan124/chirpy/internal/auth"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func (cfg *apiConfig) deleteChirp(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		log.Printf("Failed to get the bearer token")
		w.WriteHeader(401)
		return
	}

	userAuth, err := auth.ValidateJWT(token, cfg.secret)
	if err != nil {
		w.WriteHeader(401)
		return
	}

	chirp_id := r.PathValue("chirp_id")
	id, err := uuid.Parse(chirp_id)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	chir, err := cfg.Queries.GetChirp(r.Context(), id)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	if chir.UserID != userAuth {
		w.WriteHeader(403)
		return
	}
	err = cfg.Queries.DeleteChirp(r.Context(), chir.ID)
	if err != nil {
		log.Printf("Failed to delete chirp")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)

}
