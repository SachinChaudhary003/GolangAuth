package jwt

import (
	"log"

	"github.com/dgrijalva/jwt-go"
)

var SECRET_KEY = []byte("gosecretkey")

func GenerateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		log.Fatal(err)
	}
	return tokenString, nil
}
