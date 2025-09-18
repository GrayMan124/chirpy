package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) middleWareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileServerHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) metrics(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	output := fmt.Sprintf("<html><body><h1>Welcome, Chirpy Admin</h1><p>Chirpy has been visited %d times!</p></body></html>", cfg.fileServerHits.Load())
	w.Write([]byte(output))
	w.Header().Set("Content-Type", "text/html")
}

func (cfg *apiConfig) reset(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	cfg.fileServerHits.Swap(0)
	if cfg.Platform != "dev" {
		w.WriteHeader(403)

	} else {
		cfg.Queries.Reset(r.Context())
	}
}
