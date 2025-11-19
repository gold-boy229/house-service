package auth

import "golang.org/x/crypto/bcrypt"

func GetHash(password string) (string, error) {
	bslice, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bslice), nil
}

func CheckPasswordCorrectness(passwordHash []byte, password []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(passwordHash, password)
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
