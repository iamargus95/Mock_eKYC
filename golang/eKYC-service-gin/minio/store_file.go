package minio

import (
	"context"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
)

func StoreFile(file *multipart.FileHeader) error {

	data, _ := file.Open()
	defer data.Close()

	name := file.Filename
	mc := GetMinio()

	size := file.Size
	_, err := mc.PutObject(context.Background(), "clients", name, data, size,
		minio.PutObjectOptions{ContentType: "application/octet-stream"})

	return err
}
