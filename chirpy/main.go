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
}

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)

	cfg := apiConfig{}
	cfg.Queries = dbQueries
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
	serveMux.Handle("POST /api/validate_chirp", http.HandlerFunc(validation))
	server.ListenAndServe()

}
