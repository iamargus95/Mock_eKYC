package v1resources

import "github.com/golang-jwt/jwt"

type SignupPayload struct {
	Name  string `binding:"required"`
	Email string `binding:"required,email"`
	Plan  string `binding:"required,oneof=basic advanced enterprise"`
}

type AuthCustomClaims struct {
	User string
	jwt.StandardClaims
}

type ImagePayload struct {
	Type string `binding:"required,oneof=face id_card"`
	File string
}
