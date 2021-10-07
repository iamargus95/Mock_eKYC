package v1service

import (
	"iamargus95/eKYC-service-gin/conn"
	authtoken "iamargus95/eKYC-service-gin/jwt"
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

	dbtranx := db.Create(&newClient)
	if dbtranx.Error != nil {
		return dbtranx.Error
	}
	db.Save(&newClient)
	return nil
}

func Image(name string, file multipart.File, filedata *multipart.FileHeader, fileType v1r.ImagePayload) error {

	var client models.Client
	var newFile models.FileUpload

	db := conn.GetDB()

	dbtranx := db.Table("clients").Select("*").Where("name = ?", name).Scan(&client)
	if dbtranx.Error != nil {
		return dbtranx.Error
	}

	minio.StoreFile(filedata)

	newFile = models.FileUpload{
		ClientID: client.ID,
		Type:     fileType.Type,
		Size:     int64(filedata.Size),
	}

	dbtranx = db.Create(&newFile)
	if dbtranx.Error != nil {
		return dbtranx.Error
	}

	db.Save(&newFile)
	return dbtranx.Error
}
