package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (cfg *apiConfig) addUser(w http.ResponseWriter, r *http.Request) {
	type user struct {
		Email string `json:"email"`
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

	us, err := cfg.Queries.CreateUser(r.Context(), usr.Email)
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
