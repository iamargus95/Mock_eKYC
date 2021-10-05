package v1service

import (
	"iamargus95/eKYC-service-gin/conn"
	authtoken "iamargus95/eKYC-service-gin/middlewares/jwt"
	"iamargus95/eKYC-service-gin/v1/models"
	v1r "iamargus95/eKYC-service-gin/v1/resources"
	"os"
)

func Signup(body v1r.SignupPayload) error {

	var newClient models.Client

	accessKey := authtoken.JWTService().GenerateToken(body.Name)
	secretKey := os.Getenv("MYSIGNINGKEY")

	db := conn.GetDB()

	newClient = models.Client{
		Name:  body.Name,
		Email: body.Email,
		Plan: models.Plan{
			Plan: body.Plan,
		},
		SecretKey: models.SecretKey{
			Accesskey: accessKey,
			Secretkey: secretKey,
		},
	}

	err := db.Create(&newClient)
	if err.Error != nil {
		return err.Error
	}
	db.Save(&newClient)
	return err.Error
}

func Image(email string) string {
	return email
}
