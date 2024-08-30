package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dagregi/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) createFeedHandler(w http.ResponseWriter, req *http.Request, user database.User) {
	type feedParams struct {
		Name string
		URL  string
	}

	decoder := json.NewDecoder(req.Body)
	params := feedParams{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	feed, err := cfg.DB.CreateFeed(req.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}

	feedFollow, err := cfg.DB.CreateFeedFollow(req.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		FeedID:    feed.ID,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't follow feed")
		return
	}

	respondWithJSON(w, http.StatusCreated, struct {
		Feed       Feed       `json:"feed"`
		FeedFollow FeedFollow `json:"feed_follow"`
	}{
		Feed:       dbFeedToFeed(feed),
		FeedFollow: dbFeedFollowToFeedFollow(feedFollow),
	})
}

func (cfg *apiConfig) getFeedsHandler(w http.ResponseWriter, req *http.Request) {
	feeds, err := cfg.DB.GetFeeds(req.Context())
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get feeds")
		return
	}

	respondWithJSON(w, http.StatusOK, dbFeedsToFeeds(feeds))
}
