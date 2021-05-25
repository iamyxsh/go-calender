package jwt

import (
	"calendly/constants/configs"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func CreateToken(userid string) (string, error) {
	claims := jwt.MapClaims{}

	claims["userid"] = userid
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := at.SignedString([]byte(configs.JWT_SECRET))

	return token, err
}

func CheckToken(c *fiber.Ctx) (string, error) {
	token := strings.Split(c.Get("Authorization"), " ")[1]
	if len(token) == 0 {
		return "", errors.New("invalid token")
	}

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signin method - %v", token.Header["alg"])
		}

		return []byte(configs.JWT_SECRET), nil
	})
	if err != nil {
		return "", errors.New("access denied")
	}

	if _, ok := t.Claims.(jwt.Claims); !ok && !t.Valid {
		return "", errors.New("access denied")
	}

	claims, ok := t.Claims.(jwt.MapClaims)

	if ok && t.Valid {
		userid, ok := claims["userid"]

		if !ok {
			return "", err
		}

		return userid.(string), nil

	}

	return "", err
}
