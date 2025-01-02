package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/daniela2001-png/rss_aggregator_project/internal/database"
	"github.com/google/uuid"
)

// The struct that we expect from the client side
type requestBody struct {
	Name string `json:"name"`
}

// HTTP handler signature for go standard library
// Creates a new user into users database
func (api *apiConf) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	params := requestBody{}
	// Convert JSON that comes from the request body and parse into a Go value called "params"
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		setErrorResponse(w, 400, "error parsing JSON body")
		return
	}
	// Create a new user into our users database
	newUser, err := api.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	})
	if err != nil {
		setErrorResponse(w, 400, "error creating a new user")
		return
	}
	// transform the response message from user database to a resposne user
	response := ConvertDataBaseUserToResponseUser(newUser)
	// send a success response
	setJSONResponse(w, 201, response)
}

// Gets a user from his/her api key
func (api *apiConf) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	response := ConvertDataBaseUserToResponseUser(user)
	setJSONResponse(w, 200, response)
}

func (api *apiConf) handlerGetNewPostsFromUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := api.DB.GetPostsByUserID(r.Context(), database.GetPostsByUserIDParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		setErrorResponse(w, 500, "error getting posts")
		return
	}
	setJSONResponse(w, 200, ConvertGetPostsByUserIDRowToSliceOfPosts(posts))
}
