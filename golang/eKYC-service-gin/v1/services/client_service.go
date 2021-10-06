package v1service

import (
	"iamargus95/eKYC-service-gin/conn"
	authtoken "iamargus95/eKYC-service-gin/middlewares/jwt"
	"iamargus95/eKYC-service-gin/middlewares/minio"
	"iamargus95/eKYC-service-gin/v1/models"
	v1r "iamargus95/eKYC-service-gin/v1/resources"
)

func Signup(body v1r.SignupPayload) error {

	var newClient models.Client

	accessKey := authtoken.JWTService().GenerateToken(body.Name)

	db := conn.GetDB()

	newClient = models.Client{
		Name:  body.Name,
		Email: body.Email,
		Plan: models.Plan{
			Plan: body.Plan,
		},
		SecretKey: models.SecretKey{
			Accesskey: accessKey,
		},
	}

	err := db.Create(&newClient)
	if err.Error != nil {
		return err.Error
	}
	db.Save(&newClient)
	return err.Error
}

func Image(email string, body v1r.ImagePayload) (string, error) {

	var newClient models.Client

	uuid, link := minio.StoreFile(body.File)

	db := conn.GetDB()

	newClient = models.Client{
		FileUpload: models.FileUpload{
			Type:       body.Type,
			BucketLink: link,
			Size:       int64(body.File.Size),
		},
	}

	err := db.Create(&newClient)
	if err != nil {
		return uuid, err.Error
	}

	db.Save(&newClient)
	return uuid, err.Error
}
