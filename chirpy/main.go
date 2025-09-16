package main

import (
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileServerHits atomic.Int32
}

func main() {
	serveMux := http.NewServeMux()
	server := http.Server{
		Handler: serveMux,
		Addr:    ":8080",
	}
	cfg := apiConfig{}
	fileSys := cfg.middleWareMetricsInc(http.FileServer(http.Dir(".")))
	strip := http.StripPrefix("/app", fileSys)
	serveMux.Handle("/app/", strip)
	serveMux.Handle("GET /api/healthz", http.HandlerFunc(readiness))
	serveMux.Handle("GET /admin/metrics", http.HandlerFunc(cfg.metrics))
	serveMux.Handle("POST /admin/reset", http.HandlerFunc(cfg.reset))
	serveMux.Handle("POST /api/validate_chirp", http.HandlerFunc(validation))
	server.ListenAndServe()

}
