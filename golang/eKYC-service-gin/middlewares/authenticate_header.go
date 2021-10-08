package middlewares

import (
	"net/http"
	"strings"

	authtoken "iamargus95/eKYC-service-gin/jwt"

	"github.com/gin-gonic/gin"
)

func EnsureLoggedIn(authtoken authtoken.JWTInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("Authorization")
		if auth == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Error": "No Authorization Header found.",
			})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(auth, "Bearer")
		tokenString = strings.TrimSpace(tokenString)
		if tokenString == auth {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Error": "Could not find bearer token.",
			})
			ctx.Abort()
			return
		}

		token, err := authtoken.ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Error": err,
			})
			ctx.Abort()
			return
		}

		name, err := authtoken.ParseToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Error": err,
			})
			ctx.Abort()
			return
		}

		ctx.JSON(http.StatusOK, gin.H{})
		ctx.Set("client_name", name)
		ctx.Next()
	}
}
