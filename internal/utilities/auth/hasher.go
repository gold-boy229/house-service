package auth

import "golang.org/x/crypto/bcrypt"

func GetHash(password string) (string, error) {
	bslice, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bslice), nil
}
