package controllers

import (
	v1resources "iamargus95/eKYC-service-gin/resources/api/v1"
	v1s "iamargus95/eKYC-service-gin/services/api/v1"

	"github.com/gin-gonic/gin"
)

type SignupController interface {
	Save(ctx *gin.Context)
}

type controller struct {
	service v1s.SignupService
}

func New(service v1s.SignupService) SignupController {
	return &controller{
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) {

	var body v1resources.Request
	ctx.BindJSON(&body)
	c.service.Save(body)
}
