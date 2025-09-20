package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	claimsStruct := jwt.RegisteredClaims{}
	tok, err := jwt.ParseWithClaims(tokenString, &claimsStruct, func(t *jwt.Token) (interface{}, error) { return []byte(tokenSecret), nil })
	if err != nil {
		// invalid token or it expired
		return uuid.Nil, err
	}
	userStrId, err := tok.Claims.GetSubject()
	if err != nil {
		// something went wrong KEKW
		return uuid.Nil, err
	}

	userId, err := uuid.Parse(userStrId)
	return userId, nil

}
