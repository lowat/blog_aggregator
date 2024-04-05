package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lowat/blog_aggregator/internal/database"
)

func (cfg *apiConfig) handlerFeedFollowsCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Feed_ID uuid.UUID
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not decode request body")
		return
	}

	feed_follow, err := cfg.DB.CreateFeed_Follow(r.Context(), database.CreateFeed_FollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.Feed_ID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not create feed follow")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedFollowsToFeedFollows(feed_follow))
}
