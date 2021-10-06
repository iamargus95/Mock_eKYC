package minio

import (
	"context"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
)

func StoreFile(file *multipart.FileHeader) (string, string) {

	data, _ := file.Open()
	defer data.Close()

	mc := GetMinio()
	size := file.Size
	uploadInfo, _ := mc.PutObject(context.Background(), "client", "image", data, size,
		minio.PutObjectOptions{ContentType: "contentType"})

	uuid := uploadInfo.ETag
	link := uploadInfo.Location
	return uuid, link
}
