package main

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/lowat/blog_aggregator/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
}

type Feed struct {
	ID            uuid.UUID  `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	Name          string     `json:"name"`
	Url           string     `json:"url"`
	UserID        uuid.UUID  `json:"user_id"`
	LastFetchedAt *time.Time `json:"last_fetched_at"`
}

func databaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:            feed.ID,
		CreatedAt:     feed.CreatedAt,
		UpdatedAt:     feed.UpdatedAt,
		Name:          feed.Name,
		Url:           feed.Url,
		UserID:        feed.UserID,
		LastFetchedAt: nullTimeToTimePtr(feed.LastFetchedAt),
	}
}

func databaseFeedsToFeeds(feeds []database.Feed) []Feed {
	result := make([]Feed, len(feeds))
	for i, feed := range feeds {
		result[i] = databaseFeedToFeed(feed)
	}
	return result
}

type FeedFollows struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowsToFeedFollows(feedFollow database.FeedFollow) FeedFollows {
	return FeedFollows{
		ID:        feedFollow.ID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
		UserID:    feedFollow.UserID,
		FeedID:    feedFollow.FeedID,
	}
}

func databaseFeedFollowsToFeedFollowsArr(feedfollows []database.FeedFollow) []FeedFollows {
	result := make([]FeedFollows, len(feedfollows))
	for i, feedfollow := range feedfollows {
		result[i] = databaseFeedFollowsToFeedFollows(feedfollow)
	}
	return result
}

type Feed_FeedFollow struct {
	Feed        Feed        `json:"feed"`
	FeedFollows FeedFollows `json:"feed_follow"`
}

func databaseFeed_FeedFollowsToFeed_FeedFollows(feed database.Feed, feedFollow database.FeedFollow) Feed_FeedFollow {
	return Feed_FeedFollow{
		Feed:        databaseFeedToFeed(feed),
		FeedFollows: databaseFeedFollowsToFeedFollows(feedFollow),
	}
}

func nullTimeToTimePtr(t sql.NullTime) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}

type Post struct {
	ID          uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Title       string     `json:"title"`
	Url         string     `json:"url"`
	Description *string    `json:"description"`
	PublishedAt *time.Time `json:"published_at"`
	FeedID      uuid.UUID  `json:"feed_id"`
}

func databasePostToPost(post database.Post) Post {
	return Post{
		ID:          post.ID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Title:       post.Title,
		Url:         post.Url,
		Description: nullStringToStringPtr(post.Description),
		PublishedAt: nullTimeToTimePtr(post.PublishedAt),
		FeedID:      post.FeedID,
	}
}

func databasePostsToPosts(posts []database.Post) []Post {
	result := make([]Post, len(posts))
	for i, post := range posts {
		result[i] = databasePostToPost(post)
	}
	return result
}

func nullStringToStringPtr(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	}
	return nil
}
