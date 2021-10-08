package middlewares

import (
	"net/http"
	"strings"

	authtoken "iamargus95/eKYC-service-gin/jwt"

	"github.com/gin-gonic/gin"
)

// Authenticates the use the API endpoint /api/v1/image

func Authenticate(ctx *gin.Context) string {

	var nilString string
	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Error": "No Authorization Header found.",
		})
		ctx.Abort()
		return nilString
	}

	tokenString := strings.TrimPrefix(auth, "Bearer")
	tokenString = strings.TrimSpace(tokenString)
	if tokenString == auth {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Error": "Could not find bearer token.",
		})
		ctx.Abort()
		return nilString
	}

	token, err := authtoken.JWTService().ValidateToken(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Error": err.Error(),
		})
		ctx.Abort()
		return nilString
	}

	name, err := authtoken.JWTService().ParseToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Error": err.Error(),
		})
		ctx.Abort()
		return nilString
	}

	return name
}
