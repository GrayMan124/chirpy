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
	hits := fmt.Sprintf("Hits: %v", cfg.fileServerHits.Load())
	w.Write([]byte(hits))
}

func (cfg *apiConfig) reset(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	cfg.fileServerHits.Swap(0)
}
