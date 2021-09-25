package main

import (
	"iamargus95/eKYC-service-gin/conn"
	"iamargus95/eKYC-service-gin/models"
	"iamargus95/eKYC-service-gin/routes"
	"log"
)

func main() {

	db := conn.GetDB()
	err1 := db.DropTableIfExists(&models.Plan{}, &models.Client{})
	if err1.Error == nil {
		err2 := db.CreateTable(&models.Client{}, &models.Plan{})
		if err2.Error != nil {
			log.Fatalf("Table creation failed.")
		}
	}
	routes.StartGin()
	defer db.Close()
}
