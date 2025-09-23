package main

import (
	"database/sql"
	"encoding/json"
	"github.com/GrayMan124/chirpy/internal/auth"
	"log"
	"net/http"
	"time"
)

func (cfg *apiConfig) RefreshToken(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		log.Fatal("Failed to get the bearer token")
		w.WriteHeader(500)
		return
	}

	refreshToken, err := cfg.Queries.GetRefreshToken(r.Context(), token)
	if err == sql.ErrNoRows {
		w.WriteHeader(401)
		log.Printf("The token does not exist")
		return
	} else if err != nil {
		w.WriteHeader(500)
		log.Printf("Failed to send select to DB")
		return
	}
	if !refreshToken.ExpiresAt.Valid || refreshToken.RevokedAt.Valid ||
		refreshToken.ExpiresAt.Time.Before(time.Now()) {
		w.WriteHeader(401)
		log.Printf("Token has expired or has been revoked")
		return
	}
	newToken, err := auth.MakeJWT(refreshToken.UserID, cfg.secret, time.Duration(1)*time.Hour)
	if err != nil {
		log.Printf("Failed to make new token")
		w.WriteHeader(500)
		return
	}
	type responseToken struct {
		Token string `json:"token"`
	}
	tok := responseToken{Token: newToken}
	data, err := json.Marshal(tok)
	w.Write(data)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
}
