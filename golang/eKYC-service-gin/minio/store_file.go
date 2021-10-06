package minio

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
)

func StoreFile(file *multipart.FileHeader) (string, string) {

	data, _ := file.Open()
	defer data.Close()

	name := file.Filename
	mc := GetMinio()

	size := file.Size
	uploadInfo, err := mc.PutObject(context.Background(), "clients", name, data, size,
		minio.PutObjectOptions{ContentType: "application/octet-stream"})

	fmt.Println(err)
	uuid := uploadInfo.ETag
	link := uploadInfo.Location
	return uuid, link
}
