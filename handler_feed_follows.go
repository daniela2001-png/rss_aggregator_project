package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/daniela2001-png/rss_aggregator_project/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/lib/pq" // pq is a pure Go Postgres driver for the database/sql package.
)

// The struct that we expect from the client side
type requestBodyFeedFollow struct {
	FeedID uuid.UUID `json:"feed_id"`
}

// HTTP handler signature for go standard library
// Creates a new relationship between an user_id and feed_id into feed_follows table
// This handler allows that an authenticaded user can follow the feeds that wants
func (api *apiConf) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	params := requestBodyFeedFollow{}
	// Convert JSON that comes from the request body and parse into a Go value called "params"
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		setErrorResponse(w, 400, "error parsing JSON body")
		return
	}
	// Create a new user into our users database
	newUserFeed, err := api.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			setErrorResponse(w, 400, fmt.Sprintf("error: user_id and feed_id must be unique, %v", err.Code.Name()))
			return
		}
		setErrorResponse(w, 400, "error creating the feed to follow")
		return
	}
	// send a success response
	setJSONResponse(w, 201, ConvertDataBaseFeedToFollowToResponseFeedToFollow(newUserFeed))
}

// Gets all the feeds associated with an user_id
func (api *apiConf) handlerGetListOfFeedsOfAnUser(w http.ResponseWriter, r *http.Request, user database.User) {
	feedsToFollow, err := api.DB.GetFeedsFollow(r.Context(), user.ID)
	if err != nil {
		setErrorResponse(w, 400, "error getting feeds")
	}

	setJSONResponse(w, 201, ConvertDataBaseListOfFeedsToFollowToResponseFeedsToFollow(feedsToFollow))
}

// Deletes a feed_id from an user_id, in that way an user_id can unfollow a given feed.
func (api *apiConf) handlerUnFollowFeedID(w http.ResponseWriter, r *http.Request, user database.User) {
	// Here  we use a URL para, cuz we are changing the state of an user, in this case, when he/she wants to unfollow a  given feed_id
	URLParamFeedID := chi.URLParam(r, "feedFollowID")
	if URLParamFeedID == "" {
		setErrorResponse(w, 400, "error you must set correctly the feed_id to unfollow")
		return
	}
	URLParamFeedIDParsed, err := uuid.Parse(URLParamFeedID)
	if err != nil {
		setErrorResponse(w, 500, "error can not convert string to uuid type")
		return
	}
	errDB := api.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		ID:     URLParamFeedIDParsed,
	})
	if errDB != nil {
		setErrorResponse(w, 400, "error deleting feed ID")
	}
	setJSONResponse(w, 200, struct{}{})
}
