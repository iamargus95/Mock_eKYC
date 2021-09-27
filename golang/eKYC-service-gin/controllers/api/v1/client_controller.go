package controllers

import (
	v1r "iamargus95/eKYC-service-gin/resources/api/v1"
	v1s "iamargus95/eKYC-service-gin/services/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(ctx *gin.Context) {

	var body v1r.Request
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": err.Error(),
		})
	} else {
		v1s.Signup(body)
		ctx.JSON(http.StatusOK, gin.H{
			"accessKey": "10-char-JWT-Token",
			"secretKey": "20-char-JWT-Token",
		})
	}

}
