package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dagregi/blog-aggregator/internal/auth"
	"github.com/dagregi/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) createUserHandler(w http.ResponseWriter, req *http.Request) {
	type userParams struct {
		Name string
	}

	decoder := json.NewDecoder(req.Body)
	params := userParams{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error decoding parameters")
		return
	}

	user, err := cfg.DB.CreateUser(req.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	respondWithJSON(w, http.StatusOK, dbUserToUser(user))
}

func (cfg *apiConfig) getUserHandler(w http.ResponseWriter, req *http.Request) {
	apiKey, err := auth.GetAPIKey(req.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find api key")
		return
	}

	user, err := cfg.DB.GetUserByAPIKey(req.Context(), apiKey)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get user")
		return
	}

	respondWithJSON(w, http.StatusOK, dbUserToUser(user))
}
