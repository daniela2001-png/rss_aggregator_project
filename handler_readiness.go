package main

import (
	"net/http"
)

// HTTP handler signature for go standard library
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	setJSONResponse(w, 200, struct{}{})

}
