package main

import (
	"encoding/json"
	"github.com/GrayMan124/chirpy/internal/chirp"
	// "github.com/GrayMan124/chirpy/internal/database"
	"log"
	"net/http"
)

func (cfg *apiConfig) apiGetChirps(w http.ResponseWriter, r *http.Request) {
	chirps, err := cfg.Queries.GetChirps(r.Context())
	if err != nil {
		w.WriteHeader(507)
		return
	}
	var outChirps []chirp.Chirp
	for _, chi := range chirps {
		outChirps = append(outChirps, chirp.Chirp(chi))
	}

	out, err := json.Marshal(outChirps)
	if err != nil {
		log.Fatal("failed marshalling")
		w.WriteHeader(508)
		return
	}

	w.WriteHeader(200)
	w.Write(out)
	w.Header().Set("Content-Type", "application/json")

}
