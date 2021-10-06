package v1controller

import (
	authtoken "iamargus95/eKYC-service-gin/middlewares/jwt"
	v1r "iamargus95/eKYC-service-gin/v1/resources"
	v1s "iamargus95/eKYC-service-gin/v1/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Signup(ctx *gin.Context) {

	var body v1r.SignupPayload
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		ctx.Abort()
		return
	}

	err = v1s.Signup(body)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"Error": err.Error(),
		})
		ctx.Abort()
		return
	}

	aKey := authtoken.JWTService().GenerateToken(body.Name)
	ctx.JSON(http.StatusOK, gin.H{
		"access_key": aKey,
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

	token, err := authtoken.JWTService().ValidateToken(tokenString)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"Error": err.Error(),
		})
		ctx.Abort()
		return
	}

	email, err := authtoken.JWTService().ParseToken(token)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"Error": err.Error(),
		})
		ctx.Abort()
		return
	}

	var body v1r.ImagePayload
	err = ctx.Bind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		ctx.Abort()
		return
	}

	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		ctx.Abort()
		return
	}

	uuid, err := v1s.Image(email, file, header, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"imageID": uuid,
	})
}
