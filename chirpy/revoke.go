package main

import (
	"github.com/GrayMan124/chirpy/internal/auth"
	"log"
	"net/http"
)

func (cfg *apiConfig) RevokeToken(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		log.Fatal("Failed to get the bearer token")
		w.WriteHeader(500)
		return
	}
	err = cfg.Queries.RevokeToken(r.Context(), token)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Failed to send select to DB")
		return
	}
	w.WriteHeader(204)
}
