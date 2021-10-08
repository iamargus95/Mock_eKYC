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

	if fileType == "id_card" {
		fileType = "idcard"
	}

	// PutObject uploads file to minio directly from filesystem.
	// PutObject(context.Context(), bucket_name, object_name, file, file_size)
	// Here file is of type os.Open(file from file_system)
	_, err := mc.PutObject(context.Background(), fileType, uuid, data, size,
		minio.PutObjectOptions{ContentType: "application/octet-stream", UserTags: map[string]string{"client": client}})

	if err != nil {
		return Nil
	}

	return imageId
}
