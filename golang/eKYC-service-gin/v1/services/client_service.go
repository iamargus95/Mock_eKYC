package v1service

import (
	"iamargus95/eKYC-service-gin/conn"
	authtoken "iamargus95/eKYC-service-gin/middlewares/jwt"
	"iamargus95/eKYC-service-gin/minio"
	"iamargus95/eKYC-service-gin/v1/models"
	v1r "iamargus95/eKYC-service-gin/v1/resources"
	"mime/multipart"
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

func Image(email string, file multipart.File, filedata *multipart.FileHeader, fileType v1r.ImagePayload) (string, error) {

	var newFile models.FileUpload
	var client models.Client

	uuid, link := minio.StoreFile(filedata)

	db := conn.GetDB()

	db.Find(&models.Client{}).Where("email = ?", email).Scan(&client)

	newFile = models.FileUpload{
		ClientID:   client.ID,
		Type:       fileType.Type,
		BucketLink: link,
		Size:       int64(filedata.Size),
	}

	err := db.Create(&newFile)
	if err != nil {
		return uuid, err.Error
	}

	db.Save(&newFile)
	return uuid, err.Error
}
