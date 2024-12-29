package auth

import (
	"errors"
	"net/http"
	"strings"
)

/*
GetAPIKey gets the API key from a request header
as follows:

	Authorization: ApiKey {insert value here}
*/
func GetAPIKey(headers http.Header) (string, error) {
	apiKey := headers.Get("Authorization")
	if len(apiKey) == 0 {
		return "", errors.New("the api key value is required")
	}
	splitHeaderValue := strings.Split(apiKey, " ")
	if len(splitHeaderValue) != 2 {
		return "", errors.New("malformed auth header value should be: Authorization: ApiKey {value}")
	}
	if splitHeaderValue[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header value")
	}
	return splitHeaderValue[1], nil
}
