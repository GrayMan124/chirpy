package main

import (
	"encoding/json"
	"github.com/GrayMan124/chirpy/internal/auth"
	"github.com/GrayMan124/chirpy/internal/database"
	"log"
	"net/http"
)

func (cfg *apiConfig) addUser(w http.ResponseWriter, r *http.Request) {
	type user struct {
		Password string `json:password`
		Email    string `json:"email"`
	}
	decoder := json.NewDecoder(r.Body)
	usr := user{}
	err := decoder.Decode(&usr)
	if err != nil {
		output := errorResponse{
			Error: "Failed to decode user",
		}
		w.WriteHeader(500)
		out, err := json.Marshal(output)
		if err != nil {
			log.Fatal("failed marshalling")
		}
		w.Write(out)
		w.Header().Set("Content-Type", "application/json")
		return
	}
	hashed_password, err := auth.HashPassword(usr.Password)
	if err != nil {
		output := errorResponse{
			Error: "Failed to hash passowrd",
		}
		w.WriteHeader(501)
		out, err := json.Marshal(output)
		if err != nil {
			log.Fatal("failed marshalling")
		}
		w.Write(out)
		w.Header().Set("Content-Type", "application/json")
		return
	}

	us, err := cfg.Queries.CreateUser(r.Context(), database.CreateUserParams{Email: usr.Email, HashedPassword: hashed_password})
	createdUser := User{
		ID:        us.ID,
		CreatedAt: us.CreatedAt,
		UpdatedAt: us.UpdatedAt,
		Email:     us.Email,
	}

	w.WriteHeader(201)
	out, err := json.Marshal(createdUser)
	if err != nil {
		log.Fatal("failed marshalling")
	}
	w.Write(out)
	w.Header().Set("Content-Type", "application/json")

}
