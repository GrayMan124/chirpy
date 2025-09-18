package main

import (
	"encoding/json"
	"github.com/GrayMan124/chirpy/internal/chirp"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func (cfg *apiConfig) SendChirp(w http.ResponseWriter, r *http.Request) {
	type chirp struct {
		Body   string    `json:"body"`
		UserId uuid.UUID `json:"user_id"`
	}
	decoder := json.NewDecoder(r.Body)
	c := chirp{}
	err := decoder.Decode(&c)
	if err != nil {
		w.WriteHeader(505)
		return
	}

	chirpBody, valid := ValidateChirp(c.Body)
	if !valid {
		w.WriteHeader(505)
		return
	}

	output := validResponse{
		CleanedBody: replaceWords(c.Body),
	}
	w.WriteHeader(200)
	out, err := json.Marshal(output)
	if err != nil {
		log.Fatal("failed marshalling")
	}
	w.Write(out)
	w.Header().Set("Content-Type", "application/json")

}
