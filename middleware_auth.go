package main

import (
	"database/sql"
	"net/http"

	"github.com/daniela2001-png/rss_aggregator_project/internal/auth"
	"github.com/daniela2001-png/rss_aggregator_project/internal/database"
)

// auth handler signature for authenticated users
type authenticated func(w http.ResponseWriter, r *http.Request, user database.User)

// This piece of code is going to be executed before we want to, for example get user information.
// For can authenticate correctly all users
func (apiCnf *apiConf) middlewareAuth(handler authenticated) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			setErrorResponse(w, 403, "auth error: "+err.Error())
			return
		}
		user, err := apiCnf.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			if err == sql.ErrNoRows {
				setErrorResponse(w, 404, "user not found")
				return
			}
			setErrorResponse(w, 400, "could not get user: "+err.Error())
			return
		}
		// Finally call the respected handler, if authentication was success
		handler(w, r, user)
	}
}
