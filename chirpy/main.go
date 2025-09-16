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
	serveMux.Handle("GET /healthz", http.HandlerFunc(readiness))
	serveMux.Handle("GET /metrics", http.HandlerFunc(cfg.metrics))
	serveMux.Handle("POST /reset", http.HandlerFunc(cfg.reset))
	server.ListenAndServe()

}
