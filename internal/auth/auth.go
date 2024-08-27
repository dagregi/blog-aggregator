package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(h http.Header) (string, error) {
	auth := h.Get("Authorization")
	if auth == "" {
		return "", errors.New("No authorization header found")
	}

	authVals := strings.Split(auth, " ")
	if authVals[0] != "ApiKey" {
		return "", errors.New("Malformed authorization header")
	}

	return authVals[1], nil
}
