package v1service

import (
	"iamargus95/eKYC-service-gin/conn"
	"iamargus95/eKYC-service-gin/models"
	"log"
)

func NewClient(clients models.Client) {

	db := conn.GetDB()
	err := db.Create(&clients)
	if err.Error != nil {
		log.Fatalf("POST operation failed.")
	}
}
