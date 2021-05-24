package jwt

import (
	"calendly/constants/configs"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userid string) (string, error) {
	claims := jwt.MapClaims{}

	claims["userid"] = userid
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return at.SignedString(configs.JWT_SECRET)
}
