package minio

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioclient *minio.Client

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	endpoint := os.Getenv("MINIOENDPOINT")
	accessKeyID := os.Getenv("MINIOACCESS")
	secretAccessKey := os.Getenv("MINIOSECRET")

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

// Creates a minio session
func GetMinio() *minio.Client {
	return minioclient
}
