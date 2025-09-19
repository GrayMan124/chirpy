package main

import (
	"database/sql"
	"github.com/GrayMan124/chirpy/internal/database"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"sync/atomic"
)

import _ "github.com/lib/pq"

type apiConfig struct {
	fileServerHits atomic.Int32
	Queries        *database.Queries
	Platform       string
}

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)

	cfg := apiConfig{}
	cfg.Queries = dbQueries
	cfg.Platform = os.Getenv("PLATFORM")

	if err != nil {
		log.Fatal("Failed to connect to DB")
	}
	serveMux := http.NewServeMux()
	server := http.Server{
		Handler: serveMux,
		Addr:    ":8080",
	}
	fileSys := cfg.middleWareMetricsInc(http.FileServer(http.Dir(".")))
	strip := http.StripPrefix("/app", fileSys)
	serveMux.Handle("/app/", strip)
	serveMux.Handle("GET /api/healthz", http.HandlerFunc(readiness))
	serveMux.Handle("GET /admin/metrics", http.HandlerFunc(cfg.metrics))
	serveMux.Handle("POST /admin/reset", http.HandlerFunc(cfg.reset))
	serveMux.Handle("POST /api/chirps", http.HandlerFunc(cfg.SendChirp))
	serveMux.Handle("POST /api/login", http.HandlerFunc(cfg.loginUser))
	serveMux.Handle("POST /api/users", http.HandlerFunc(cfg.addUser))
	serveMux.Handle("GET /api/chirps", http.HandlerFunc(cfg.apiGetChirps))
	serveMux.Handle("GET /api/chirps/{chirp_id}", http.HandlerFunc(cfg.apiGetChirp))
	server.ListenAndServe()
}
