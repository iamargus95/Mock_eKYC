package authtoken

import (
	"fmt"
	"os"

	v1r "iamargus95/eKYC-service-gin/v1/resources"

	"github.com/golang-jwt/jwt"
)

func IsValid(tokenString string) (string, error) {

	var user string
	var claims v1r.AuthCustomClaims
	signingKey := os.Getenv("MYSIGNINGKEY")

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if claims, ok := token.Claims.(*v1r.AuthCustomClaims); ok && token.Valid {
		user = fmt.Sprintf(claims.User)
	}
	return user, err
}
