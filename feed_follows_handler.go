package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dagregi/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) createFeedFollowHandler(
	w http.ResponseWriter,
	req *http.Request,
	user database.User,
) {
	type feedFollowParams struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(req.Body)
	params := feedFollowParams{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error decoding parameters")
		return
	}

	feedFollow, err := cfg.DB.CreateFeedFollow(req.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't follow feed")
		return
	}

	respondWithJSON(w, http.StatusOK, dbFeedFollowToFeedFollow(feedFollow))
}

func (cfg *apiConfig) deleteFeedFollowHandler(
	w http.ResponseWriter,
	req *http.Request,
	user database.User,
) {
	feedFollowID, err := uuid.Parse(req.PathValue("feedFollowID"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid feed follow id")
		return
	}

	err = cfg.DB.DeleteFeedFollow(req.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete follow")
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}

func (cfg *apiConfig) getFeedFollowsHandler(
	w http.ResponseWriter,
	req *http.Request,
	user database.User,
) {
	feedFollows, err := cfg.DB.GetFeedFollowByUserID(req.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get follows")
		return
	}

	respondWithJSON(w, http.StatusOK, dbFeedFollowsToFeedFollows(feedFollows))
}
