package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	authorization := headers.Get("Authorization")
	if authorization == "" {
		return "", fmt.Errorf("Header does not exist")
	}
	splitAuth := strings.Split(authorization, " ")
	token := splitAuth[1]
	return token, nil

}
