package v1controller

import (
	"fmt"
	authtoken "iamargus95/eKYC-service-gin/middlewares/jwt"
	v1r "iamargus95/eKYC-service-gin/v1/resources"
	v1s "iamargus95/eKYC-service-gin/v1/services"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func Signup(ctx *gin.Context) {

	var body v1r.SignupPayload
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": err.Error(),
		})
		ctx.Abort()
		return
	}

	err = v1s.Signup(body)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	aKey, _ := authtoken.GenerateJWT(body.Name)
	sKey := os.Getenv("MYSIGNINGKEY")
	ctx.JSON(http.StatusOK, gin.H{
		"accessKey": aKey,
		"secretKey": sKey,
	})
}

func Image(ctx *gin.Context) {

	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		ctx.JSON(http.StatusForbidden, gin.H{
			"Error": "No Authorization Header found.",
		})
		ctx.Abort()
		return
	}

	tokenString := strings.TrimPrefix(auth, "Bearer")
	tokenString = strings.TrimSpace(tokenString)
	if tokenString == auth {
		ctx.JSON(http.StatusForbidden, gin.H{
			"Error": "Could not find bearer token.",
		})
		ctx.Abort()
		return
	}

	token, err := authtoken.IsValid(tokenString)
	fmt.Println(token)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"Error": err.Error(),
		})
		ctx.Abort()
		return
	}

	fmt.Printf("token: %v\n", token)
}
