package minio

import (
	"context"

	"github.com/minio/minio-go/v7"
)

func GetFile(name, bucket string) (*minio.Object, error) {

	var NIL *minio.Object

	if bucket == "id_card" {
		bucket = "idcard"
	}

	mc := GetMinio()

	obj, err := mc.GetObject(context.Background(), bucket, name, minio.GetObjectOptions{})
	if err != nil {
		return NIL, err
	}

	return obj, nil
}
