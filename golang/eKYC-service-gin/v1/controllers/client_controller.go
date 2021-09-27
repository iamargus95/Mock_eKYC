package v1controller

import (
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
			"errorMessage": err.Error(),
		})
		ctx.Abort()
	} else {
		v1s.Signup(body)
		ctx.JSON(http.StatusOK, gin.H{
			"accessKey": "10-char-JWT-Token",
			"secretKey": "20-char-JWT-Token",
		})
	}

}
