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

func Image(name string, file multipart.File, filedata *multipart.FileHeader, fileType v1r.ImagePayload) (string, error) {

	var client models.Client
	var newFile models.FileUpload

	db := conn.GetDB()

	err := db.Table("clients").Select("*").Where("name = ?", name).Scan(&client)
	if err.Error != nil {
		return "Error1", err.Error //troubleshooting
	}

	uuid, link := minio.StoreFile(filedata)

	newFile = models.FileUpload{
		ClientID:   client.ID,
		Type:       fileType.Type,
		BucketLink: link,
		Size:       int64(filedata.Size),
	}

	err = db.Create(&newFile)
	if err != nil {
		return "Error2", err.Error //troubleshooting
	}

	db.Save(&newFile)
	return uuid, err.Error
}
