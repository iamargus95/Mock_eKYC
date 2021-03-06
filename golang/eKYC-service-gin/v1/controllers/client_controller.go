package v1controller

import (
	"fmt"
	authtoken "iamargus95/eKYC-service-gin/jwt"
	v1r "iamargus95/eKYC-service-gin/v1/resources"
	v1s "iamargus95/eKYC-service-gin/v1/services"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

const (
	SIZE = 5000000
)

// Signup controller helps create new clients inthe database, A successful request returns an access key necessary for further operations.
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

// Image controller helps client upload images to the database, It returns a UUID for the uploaded image.
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

	if header.Size > SIZE {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "file size should be less than 5MB",
		})
		ctx.Abort()
		return
	}

	buff := make([]byte, 512) // see http://golang.org/pkg/net/http/#DetectContentType
	_, _ = file.Read(buff)
	fileType := http.DetectContentType(buff)

	re, _ := regexp.Compile(fileType)
	matched := re.FindString("image/png image/jpeg")
	if matched == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid filetype, must be jpeg or png.",
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

// FaceMatch controller helps the clients match 2 images with each other, a match score between 0-100.
func FaceMatch(ctx *gin.Context) {

	var body v1r.FaceMatchPayload
	ctxData, _ := ctx.Get("client_name")
	client_name := fmt.Sprint(ctxData)

	err := ctx.Bind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		ctx.Abort()
		return
	}

	matchScore, err := v1s.GetMatch(client_name, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"score": matchScore,
	})
}
