package main

import (
	"time"

	"github.com/daniela2001-png/rss_aggregator_project/internal/database"
	"github.com/google/uuid"
)

type ResponseUser struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

type ResponseFeed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    string    `json:"user_id"`
}

type ResponseFeedToFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func ConvertDataBaseUserToResponseUser(dbUser database.User) ResponseUser {
	return ResponseUser{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

func ConvertDataBaseFeedToResponseFeed(dbFeed database.Feed) ResponseFeed {
	return ResponseFeed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID.String(),
	}
}

func ConvertDataBaseListOfFeedsToResponseFeeds(dbFeeds []database.Feed) []ResponseFeed {
	feeds := []ResponseFeed{}
	for _, value := range dbFeeds {
		feeds = append(feeds, ConvertDataBaseFeedToResponseFeed(value))
	}
	return feeds
}

func ConvertDataBaseFeedToFollowToResponseFeedToFollow(dbFeedFollow database.FeedFollow) ResponseFeedToFollow {
	return ResponseFeedToFollow{
		ID:        dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID:    dbFeedFollow.UserID,
		FeedID:    dbFeedFollow.FeedID,
	}
}

func ConvertDataBaseListOfFeedsToFollowToResponseFeedsToFollow(dbFeedsToFollow []database.FeedFollow) []ResponseFeedToFollow {
	feeds := []ResponseFeedToFollow{}
	for _, value := range dbFeedsToFollow {
		feeds = append(feeds, ConvertDataBaseFeedToFollowToResponseFeedToFollow(value))
	}
	return feeds
}
