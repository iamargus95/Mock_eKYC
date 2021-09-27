package main

import (
	"iamargus95/eKYC-service-gin/conn"
	"iamargus95/eKYC-service-gin/v1/models"
	"iamargus95/eKYC-service-gin/v1/routes"
)

func main() {

	db := conn.GetDB()
	defer db.Close()

	db.Debug().AutoMigrate(&models.Client{}, &models.Plan{})
	routes.StartGin()
}
