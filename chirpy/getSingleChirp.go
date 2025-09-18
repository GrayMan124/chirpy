package main

import (
	"encoding/json"
	"github.com/GrayMan124/chirpy/internal/chirp"
	// "github.com/GrayMan124/chirpy/internal/database"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func (cfg *apiConfig) apiGetChirp(w http.ResponseWriter, r *http.Request) {
	chirp_id := r.PathValue("chirp_id")
	id, err := uuid.Parse(chirp_id)
	if err != nil {
		w.WriteHeader(508)
		return
	}
	chir, err := cfg.Queries.GetChirp(r.Context(), id)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	var outChirp chirp.Chirp
	outChirp = chirp.Chirp(chir)

	out, err := json.Marshal(outChirp)
	if err != nil {
		log.Fatal("failed marshalling")
		w.WriteHeader(508)
		return
	}

	w.WriteHeader(200)
	w.Write(out)
	w.Header().Set("Content-Type", "application/json")

}
