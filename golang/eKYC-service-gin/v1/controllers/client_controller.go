package v1controller

import (
	"fmt"
	authtoken "iamargus95/eKYC-service-gin/jwt"
	v1r "iamargus95/eKYC-service-gin/v1/resources"
	v1s "iamargus95/eKYC-service-gin/v1/services"
	"net/http"

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

	var body v1r.ImagePayload
	err := ctx.Bind(&body)
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

	ctxData, _ := ctx.Get("client_name")
	client_name := fmt.Sprint(ctxData)

	uuid, err := v1s.ImageUpload(client_name, file, header, body)
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

func FaceMatch(ctx *gin.Context) {

	var body v1r.FaceMatchPayload

	err := ctx.Bind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "invalid request",
		})
	}

	ctxData, _ := ctx.Get("client_name")
	client_name := fmt.Sprint(ctxData)
	matchScore, err := v1s.GetMatch(client_name, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"score": matchScore,
	})
}
