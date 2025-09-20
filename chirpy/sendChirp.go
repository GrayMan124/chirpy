package main

import (
	"encoding/json"
	"github.com/GrayMan124/chirpy/internal/auth"
	"github.com/GrayMan124/chirpy/internal/chirp"
	"github.com/GrayMan124/chirpy/internal/database"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func (cfg *apiConfig) SendChirp(w http.ResponseWriter, r *http.Request) {
	type chirpJson struct {
		Body   string    `json:"body"`
		UserId uuid.UUID `json:"user_id"`
	}
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		log.Fatal("Failed to get the bearer token")
		return
	}
	decoder := json.NewDecoder(r.Body)
	c := chirpJson{}
	erro := decoder.Decode(&c)
	if erro != nil {
		w.WriteHeader(505)
		return
	}

	userAuth, err := auth.ValidateJWT(token, cfg.secret)
	if err != nil {
		w.WriteHeader(401)
		return
	}

	chirpBody, valid := chirp.ValidateChirp(c.Body)
	if !valid {
		w.WriteHeader(506)
		return
	}
	c.Body = chirpBody
	createdChirp, err := cfg.Queries.InsertChirp(r.Context(), database.InsertChirpParams{Body: c.Body, UserID: userAuth})
	if err != nil {
		w.WriteHeader(507)
		return
	}
	returnChirp := chirp.Chirp{
		ID:        createdChirp.ID,
		CreatedAt: createdChirp.CreatedAt,
		UpdatedAt: createdChirp.CreatedAt,
		Body:      createdChirp.Body,
		UserID:    createdChirp.UserID,
	}

	out, err := json.Marshal(returnChirp)
	if err != nil {
		log.Fatal("failed marshalling")
		w.WriteHeader(508)
		return
	}
	w.WriteHeader(201)
	w.Write(out)
	w.Header().Set("Content-Type", "application/json")

}
