package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("Authorization header not found in header")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("Authorization header does not contain valid authorization header")
	}

	return vals[1], nil
}
