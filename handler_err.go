package main

import (
	"net/http"
)

// HTTP handler signature for go standard library
func handlerError(w http.ResponseWriter, r *http.Request) {
	setErrorResponse(w, 400, "Something was wrong")
}
