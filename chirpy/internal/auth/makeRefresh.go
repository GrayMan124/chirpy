package auth

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

func MakeRefreshToken() (string, error) {
	buffer := make([]byte, 32)
	_, err := rand.Read(buffer)
	if err != nil {
		log.Printf("Failed to generate random string")
		return "", err
	}
	output := hex.EncodeToString(buffer)
	return output, err
}
