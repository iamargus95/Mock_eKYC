package v1service

import (
	"iamargus95/eKYC-service-gin/conn"
	"iamargus95/eKYC-service-gin/models"
	v1r "iamargus95/eKYC-service-gin/resources/api/v1"
)

func Signup(body v1r.Request) {

	var newClient models.Client

	db := conn.GetDB()

	newClient = models.Client{
		Name:  body.Name,
		Email: body.Email,
		Plan: models.Plan{
			Plan: body.Plan,
		},
	}

	db.Create(&newClient)
	db.Save(&newClient)
}
