package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {

	hashedPswd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hashedPswd), nil
}
