package main

import (
	"encoding/json"
	"github.com/GrayMan124/chirpy/internal/auth"
	// "github.com/GrayMan124/chirpy/internal/database"
	"log"
	"net/http"
	"time"
)

func (cfg *apiConfig) loginUser(w http.ResponseWriter, r *http.Request) {
	type user struct {
		Password string `json:"password"`
		Email    string `json:"email"`
		ExpireIn int    `json:"expires_in_seconds"`
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
	expires := min(3600, usr.ExpireIn)
	if expires == 0 {
		expires = 3600
	}
	token, err := auth.MakeJWT(usrData.ID, cfg.secret, time.Duration(expires)*time.Second)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	createdUser := User{
		ID:        usrData.ID,
		CreatedAt: usrData.CreatedAt,
		UpdatedAt: usrData.UpdatedAt,
		Email:     usrData.Email,
		Token:     token,
	}

	w.WriteHeader(200)
	out, err := json.Marshal(createdUser)
	if err != nil {
		log.Fatal("failed marshalling")
	}
	w.Write(out)
	w.Header().Set("Content-Type", "application/json")

}
