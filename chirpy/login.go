package main

import (
	"encoding/json"
	"github.com/GrayMan124/chirpy/internal/auth"
	// "github.com/GrayMan124/chirpy/internal/database"
	"log"
	"net/http"
)

func (cfg *apiConfig) loginUser(w http.ResponseWriter, r *http.Request) {
	type user struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	decoder := json.NewDecoder(r.Body)
	usr := user{}
	err := decoder.Decode(&usr)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	usrData, err := cfg.Queries.GetUsrEmail(r.Context(), usr.Email)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	err = auth.CheckPasswordHash(usr.Password, usrData.HashedPassword)
	if err != nil {
		w.WriteHeader(401)
		return
	}
	createdUser := User{
		ID:        usrData.ID,
		CreatedAt: usrData.CreatedAt,
		UpdatedAt: usrData.UpdatedAt,
		Email:     usrData.Email,
	}

	w.WriteHeader(200)
	out, err := json.Marshal(createdUser)
	if err != nil {
		log.Fatal("failed marshalling")
	}
	w.Write(out)
	w.Header().Set("Content-Type", "application/json")

}
