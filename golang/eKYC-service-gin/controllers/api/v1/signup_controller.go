package controllers

import (
	"iamargus95/eKYC-service-gin/models"
	v1s "iamargus95/eKYC-service-gin/services/api/v1"

	"github.com/gin-gonic/gin"
)

func NewClient(ctx *gin.Context) {

	var body models.Client
	err := ctx.BindJSON(&body)
	if err == nil {
		v1s.NewClient(body)
	}
}
