package authtoken

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type authCustomClaims struct {
	User   string
	IsUser bool
	jwt.StandardClaims
}

var mySigningKey = []byte(os.Getenv("MYSIGNINGKEY"))

func GenerateJWT(user string, isUser bool) (string, error) {

	claims := &authCustomClaims{
		user,
		isUser,
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
