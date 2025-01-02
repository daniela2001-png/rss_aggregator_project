package main

import (
	"log"
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

func ConvertRSSItemsListToDatabaseCreatePostParams(rssItems []RSSItem, feedID uuid.UUID) []database.CreatePostParams {
	posts := []database.CreatePostParams{}
	for _, value := range rssItems {
		pubDate, err := time.Parse(time.RFC1123, value.PubDate)
		if err != nil {
			log.Println("can not parse string to time format: ", err)
			return nil
		}
		posts = append(posts, database.CreatePostParams{
			ID:          uuid.New(),
			Title:       value.Title,
			Description: value.Description,
			Link:        value.Link,
			PubDate:     pubDate,
			FeedID:      feedID,
		})
	}
	return posts
}

type Post struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	PubDate     time.Time `json:"publication_date"`
}

func ConvertGetPostsByUserIDRowToSliceOfPosts(dbPosts []database.GetPostsByUserIDRow) []Post {
	posts := []Post{}
	for _, value := range dbPosts {
		posts = append(posts, Post{
			Title:       value.Title,
			Description: value.Description,
			Link:        value.Link,
			PubDate:     value.PubDate,
		})
	}
	return posts
}
