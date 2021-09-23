package v1service

import (
	u "iamargus95/eKYC-service-gin/apiHelpers"
	"iamargus95/eKYC-service-gin/conn"
	"iamargus95/eKYC-service-gin/models"
	v1r "iamargus95/eKYC-service-gin/resources/api/v1"
)

type SignupService interface {
	Save(v1r.Request) map[string]interface{}
}

type signupService struct {
	clients []v1r.Request
}

func New() SignupService {
	return &signupService{}
}

func (service *signupService) FindAll() map[string]interface{} {
	db := conn.GetDB()
	db.Find(&service.clients)
	resp := u.Message(200, "Success")
	return resp
}

func (service *signupService) Save(body v1r.Request) map[string]interface{} {
	//DB needed
	//Include ORM here.
	//Pass entire struct as arg

	db := conn.GetDB()
	db.Create(&models.Client{
		Name:  body.Name,
		Email: body.Email,
		Plan:  models.Plan{Plan: body.Plan},
	})
	resp := u.Message(200, "Success")
	return resp
}
