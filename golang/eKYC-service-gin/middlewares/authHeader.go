package middlewares

import (
	"net/http"
	"strings"

	authtoken "iamargus95/eKYC-service-gin/jwt"

	"github.com/gin-gonic/gin"
)

func ValidHeader(ctx *gin.Context) string {

	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		ctx.JSON(http.StatusForbidden, gin.H{
			"Error": "No Authorization Header found.",
		})
		ctx.Abort()
		return ""
	}

	tokenString := strings.TrimPrefix(auth, "Bearer")
	tokenString = strings.TrimSpace(tokenString)
	if tokenString == auth {
		ctx.JSON(http.StatusForbidden, gin.H{
			"Error": "Could not find bearer token.",
		})
		ctx.Abort()
		return ""
	}

	token, err := authtoken.JWTService().ValidateToken(tokenString)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"Error": err.Error(),
		})
		ctx.Abort()
		return ""
	}

	name, err := authtoken.JWTService().ParseToken(token)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"Error": err.Error(),
		})
		ctx.Abort()
		return ""
	}

	return name
}