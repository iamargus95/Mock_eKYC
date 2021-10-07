package minio

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioclient *minio.Client

func init() {

	endpoint := "localhost:9000"
	accessKeyID := "myaccesskey"
	secretAccessKey := "mysecretkey"

	// Initialize minio client object.
	newClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}

	minioclient = newClient
}

func GetMinio() *minio.Client {
	return minioclient
}
