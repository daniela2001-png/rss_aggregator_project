package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/daniela2001-png/rss_aggregator_project/internal/database"
	"github.com/google/uuid"
)

type requestBody struct {
	Name string `json:"name"`
}

// HTTP handler signature for go standard library
func (api *apiConf) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	params := requestBody{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		setErrorResponse(w, 400, "error parsing JSON body")
		return
	}
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
	// transform the response message
	response := ConvertDataBaseUserToResponseUser(newUser)
	setJSONResponse(w, 200, response)
}
