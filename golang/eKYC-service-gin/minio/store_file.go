package minio

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

func StoreFile(client, fileType string, file *multipart.FileHeader) uuid.UUID {

	var Nil uuid.UUID
	data, _ := file.Open()
	defer data.Close()

	mc := GetMinio()

	size := file.Size

	imageId, _ := uuid.NewUUID()
	uuid := fmt.Sprintf("%v", imageId) //Convert uuid.UUID to string

	_, err := mc.PutObject(context.Background(), fileType, uuid, data, size, //Use FileType of face or id_card as bucket names. Separate the functions and have a CreateBucket func.
		minio.PutObjectOptions{ContentType: "application/octet-stream", UserTags: map[string]string{"client": client}}) //Create tags for stored image

	if err != nil {
		return Nil
	}

	return imageId
}
