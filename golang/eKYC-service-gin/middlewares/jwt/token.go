package authtoken

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTInterface interface {
	GenerateToken(email string) string
	ValidateToken(token string) (*jwt.Token, error)
	ParseToken(token *jwt.Token) (string, error)
}

type authCustomClaims struct {
	Email string
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
}

func JWTService() JWTInterface {
	secret := os.Getenv("MYSIGNINGKEY")
	return &jwtServices{
		secretKey: secret,
	}
}

func (authtoken *jwtServices) GenerateToken(email string) string {
	claims := &authCustomClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(authtoken.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (authtoken *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(encodedToken, &authCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token: %s", token.Header["alg"])

		}
		return []byte(authtoken.secretKey), nil
	})
}

func (authtoken *jwtServices) ParseToken(token *jwt.Token) (string, error) {

	var email string
	var err error
	if claims, ok := token.Claims.(*authCustomClaims); ok && token.Valid {
		email = fmt.Sprintf(claims.Email)
	} else {
		err = fmt.Errorf("token validation failed")
	}

	return email, err
}
