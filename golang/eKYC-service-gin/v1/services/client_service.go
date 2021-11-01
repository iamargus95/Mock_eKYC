package v1service

import (
	"iamargus95/eKYC-service-gin/conn"
	authtoken "iamargus95/eKYC-service-gin/jwt"
	"iamargus95/eKYC-service-gin/minio"
	rs "iamargus95/eKYC-service-gin/randomScore"
	"iamargus95/eKYC-service-gin/v1/models"
	v1r "iamargus95/eKYC-service-gin/v1/resources"
	"mime/multipart"

	"github.com/google/uuid"
)

func Signup(body v1r.SignupPayload) error {

	var newClient models.Client

	accessKey := authtoken.JWTService().GenerateToken(body.Name)
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

	dbtranx := conn.DB.Create(&newClient)
	if dbtranx.Error != nil {
		return dbtranx.Error
	}
	conn.DB.Save(&newClient)
	return nil
}

func ImageUpload(clientName string, file multipart.File, filedata *multipart.FileHeader, fileType v1r.ImagePayload) (uuid.UUID, error) {

	var Nil uuid.UUID
	var client models.Client
	var newFile models.FileUpload

	dbtranx := conn.DB.Table("clients").Select("ID").Where("name = ?", clientName).Scan(&client)
	if dbtranx.Error != nil {
		return Nil, dbtranx.Error
	}

	uuid := minio.StoreFile(clientName, fileType.Type, filedata)

	newFile = models.FileUpload{
		ClientID: client.ID,
		Type:     fileType.Type,
		UUID:     uuid,
		Size:     int64(filedata.Size),
	}

	dbtranx = conn.DB.Create(&newFile)
	if dbtranx.Error != nil {
		return Nil, dbtranx.Error
	}

	conn.DB.Save(&newFile)
	return uuid, dbtranx.Error
}

func GetMatch(name string, body v1r.FaceMatchPayload) (int, error) {

	var client models.Client
	var upload models.FileUpload

	dbtranx := conn.DB.Table("clients").Select("ID").Where("name = ?", name).Scan(&client)
	if dbtranx.Error != nil {
		return 0, dbtranx.Error
	}

	dbtranx = conn.DB.Table("file_upload").Select("type").Where("client_id = ?", client.ID).
		Where("UUID = ?", body.Image1).Scan(&upload)
	if dbtranx.Error != nil {
		return 0, dbtranx.Error
	}

	_, err := minio.GetFile(body.Image1.String(), upload.Type)
	if err != nil {
		return 0, err
	}

	dbtranx = conn.DB.Table("file_upload").Select("type").Where("client_id = ?", client.ID).
		Where("UUID = ?", body.Image2).Scan(&upload)
	if dbtranx.Error != nil {
		return 0, dbtranx.Error
	}

	_, err = minio.GetFile(body.Image2.String(), upload.Type)
	if err != nil {
		return 0, err
	}

	score := rs.GenerateScore()

	var newScore = models.FaceMatch{
		ClientID: client.ID,
		Image1:   body.Image1,
		Image2:   body.Image2,
		Score:    uint(score),
	}

	dbtranx = conn.DB.Create(&newScore)
	if dbtranx.Error != nil {
		return 0, dbtranx.Error
	}

	conn.DB.Save(&newScore)

	return score, nil
}
