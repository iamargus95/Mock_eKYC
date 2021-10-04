package authtoken

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func IsValid(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, isvalid := t.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token %v", t.Header["alg"])
		}
		return []byte(os.Getenv("MYSIGNINGKEY")), nil
	})
}

func DecodeToken(token *jwt.Token) string {
	return "client"
}
