package v1service

import (
	"iamargus95/eKYC-service-gin/conn"
	"iamargus95/eKYC-service-gin/v1/models"
	v1r "iamargus95/eKYC-service-gin/v1/resources"
)

func Signup(body v1r.SignupPayload) {

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
