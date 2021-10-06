package minio

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioclient *minio.Client

func init() {

	endpoint := "172.19.0.3:9090"
	accessKeyID := "minio123"
	secretAccessKey := "minio456" //Doubtful

	// Initialize minio client object.
	newClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: true,
	})
	if err != nil {
		log.Fatalln(err)
	}

	minioclient = newClient
}

func GetMinio() *minio.Client {
	return minioclient
}
