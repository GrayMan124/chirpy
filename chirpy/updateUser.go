package main

import (
	"encoding/json"
	"github.com/GrayMan124/chirpy/internal/auth"
	"github.com/GrayMan124/chirpy/internal/database"
	"log"
	"net/http"
)

func (cfg *apiConfig) updateUser(w http.ResponseWriter, r *http.Request) {
	type userJson struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		log.Printf("Failed to get the bearer token")
		w.WriteHeader(401)
		return
	}
	decoder := json.NewDecoder(r.Body)
	usr := userJson{}
	err = decoder.Decode(&usr)
	if err != nil {
		w.WriteHeader(505)
		return
	}
	userAuth, err := auth.ValidateJWT(token, cfg.secret)
	if err != nil {
		w.WriteHeader(401)
		return
	}

	hashed_password, err := auth.HashPassword(usr.Password)
	updated_usr, err := cfg.Queries.UpdateUser(r.Context(),
		database.UpdateUserParams{
			Email:          usr.Email,
			HashedPassword: hashed_password,
			ID:             userAuth})

	createdUser := User{
		ID:        updated_usr.ID,
		CreatedAt: updated_usr.CreatedAt,
		UpdatedAt: updated_usr.UpdatedAt,
		Email:     updated_usr.Email,
	}

	w.WriteHeader(200)
	out, err := json.Marshal(createdUser)
	if err != nil {
		log.Fatal("failed marshalling")
	}
	w.Write(out)
	w.Header().Set("Content-Type", "application/json")

}
