package jwt

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte(os.Getenv("MYSIGNINGKEY"))

func GenerateJWT(user string) (string, error) {
	fmt.Print(mySigningKey)
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorised"] = true
	claims["user"] = user
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Fatal(err)
	}

	return tokenString, err
}
