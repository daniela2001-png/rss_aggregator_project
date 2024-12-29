package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/daniela2001-png/rss_aggregator_project/internal/database"
	"github.com/google/uuid"
)

// The struct that we expect from the client side
type requestBodyFeed struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// HTTP handler signature for go standard library
// Creates a new feed into feeds table
func (api *apiConf) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	params := requestBodyFeed{}
	// Convert JSON that comes from the request body and parse into a Go value called "params"
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		setErrorResponse(w, 400, "error parsing JSON body")
		return
	}
	// Create a new user into our users database
	newFeed, err := api.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		setErrorResponse(w, 400, "error creating a new feed")
		return
	}
	// send a success response
	setJSONResponse(w, 201, ConvertDataBaseFeedToResponseFeed(newFeed))
}

// HTTP handler signature for go standard library
// Gets a list of feeds from feeds table
func (api *apiConf) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	// Create a new user into our users database
	feeds, err := api.DB.GetFeeds(r.Context())
	if err != nil {
		setErrorResponse(w, 400, "error getting feeds")
		return
	}
	// send a success response
	setJSONResponse(w, 200, ConvertDataBaseListOfFeedsToResponseFeeds(feeds))
}
