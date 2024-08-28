package main

import (
	"net/http"
	"strconv"

	"github.com/dagregi/blog-aggregator/internal/database"
)

func (cfg *apiConfig) getPostsHandler(w http.ResponseWriter, req *http.Request, user database.User) {
	limitQuery := req.URL.Query().Get("limit")
	limit := 10
	if specifiedLimit, err := strconv.Atoi(limitQuery); err == nil {
		limit = specifiedLimit
	}

	posts, err := cfg.DB.GetPostsForUser(req.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get posts")
		return
	}

	respondWithJSON(w, http.StatusOK, dbPostsToPosts(posts))
}
