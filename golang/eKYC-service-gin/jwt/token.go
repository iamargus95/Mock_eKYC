//Package authtoken provides a way to generate, validate and parse a JWT.
package authtoken

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTInterface interface {
	GenerateToken(Name string) string
	ValidateToken(token string) (*jwt.Token, error)
	ParseToken(token *jwt.Token) (string, error)
}

type authCustomClaims struct {
	Name string
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

// Generates JWT and uses client name as input string for payload that is to be signed.
func (authtoken *jwtServices) GenerateToken(name string) string {
	claims := &authCustomClaims{
		name,
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

// Validates a given signedString against a secret and returns the JWT with just the payload data.
func (authtoken *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {

	return jwt.ParseWithClaims(encodedToken, &authCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token: %s", token.Header["alg"])
		}
		return []byte(authtoken.secretKey), nil
	})
}

// Parses a JWT and retrieves payload data from it.
// Here the payload data is the client name.
func (authtoken *jwtServices) ParseToken(token *jwt.Token) (string, error) {

	var name string
	var err error
	if claims, ok := token.Claims.(*authCustomClaims); ok && token.Valid {
		name = fmt.Sprintf(claims.Name)
	} else {
		err = fmt.Errorf("token validation failed")
	}

	return name, err
}
