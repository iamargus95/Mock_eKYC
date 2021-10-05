package authtoken

import (
	"fmt"
	"os"
	"time"

	v1r "iamargus95/eKYC-service-gin/v1/resources"

	"github.com/golang-jwt/jwt"
)

var mySigningKey = []byte(os.Getenv("MYSIGNINGKEY"))

func GenerateJWT(user string) (string, error) {

	claims := v1r.AuthCustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", fmt.Errorf("invalid signing key %v", err)
	}
	return t, err
}
