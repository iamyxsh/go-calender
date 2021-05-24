package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPass(pass string) (string, error) {
	password := []byte(pass)

	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	return string(hash), err
}
