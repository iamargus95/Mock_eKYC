package v1controller

import (
	"iamargus95/eKYC-service-gin/middlewares/jwt"
	v1r "iamargus95/eKYC-service-gin/v1/resources"
	v1s "iamargus95/eKYC-service-gin/v1/services"
	"net/http"
	"os"

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

	aKey, _ := jwt.GenerateJWT(body.Name)
	sKey := os.Getenv("MYSIGNINGKEY")
	ctx.JSON(http.StatusOK, gin.H{
		"accessKey": aKey,
		"secretKey": sKey,
	})
}
