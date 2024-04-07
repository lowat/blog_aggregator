package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lowat/blog_aggregator/internal/database"
)

func (cfg *apiConfig) handlerFeedsCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not decode request body")
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}

	feedFollow, err := cfg.DB.CreateFeed_Follow(r.Context(), database.CreateFeed_FollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed follow")
		return
	}

	type Feed_FeedFollow struct {
	}

	respondWithJSON(w, http.StatusOK, databaseFeed_FeedFollowsToFeed_FeedFollows(feed, feedFollow))
}
