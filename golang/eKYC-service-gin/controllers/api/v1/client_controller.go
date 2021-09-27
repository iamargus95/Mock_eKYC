package controllers

import (
	v1r "iamargus95/eKYC-service-gin/resources/api/v1"
	v1s "iamargus95/eKYC-service-gin/services/api/v1"
	"log"

	"github.com/gin-gonic/gin"
)

func Signup(ctx *gin.Context) {

	var body v1r.Request
	err := ctx.BindJSON(&body)
	if err != nil {
		log.Fatal(err)
	} else {
		v1s.Signup(body)
	}

}
