package main

import (
	"iamargus95/eKYC-service-gin/conn"
	"iamargus95/eKYC-service-gin/models"
	"iamargus95/eKYC-service-gin/routes"
)

func main() {

	db := conn.GetDB()
	db.Debug().AutoMigrate(&models.Client{}, &models.Plan{})
	routes.StartGin()
	defer db.Close()
}
