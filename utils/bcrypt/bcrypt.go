package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPass(pass string) (string, error) {
	password := []byte(pass)

	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	return string(hash), err
}

func CompPass(pass string, hash string) error {
	password := []byte(pass)
	hashed := []byte(hash)

	return bcrypt.CompareHashAndPassword(hashed, password)
}
