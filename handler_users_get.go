package main

import (
	"net/http"

	"github.com/lowat/blog_aggregator/internal/auth"
)

func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request) {
	api_key, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized: No APIKey provided")
	}

	user, err := cfg.DB.GetUserByAPIKey(r.Context(), api_key)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve user")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
