package main

import (
	"net/http"

	"github.com/dagregi/blog-aggregator/internal/auth"
	"github.com/dagregi/blog-aggregator/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
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

		handler(w, req, user)
	}
}
