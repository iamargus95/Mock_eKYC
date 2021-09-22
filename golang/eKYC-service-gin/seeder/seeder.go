package seed

import (
	"iamargus95/eKYC-service-gin/conn"
	"iamargus95/eKYC-service-gin/models"
	"log"
)

var clients = []models.Client{
	{
		Name:  "E Corp",
		Email: "ecorp@evil.com",
		Plan:  models.Plan{ClientID: 1, Plan: "Enterprise"},
	},
	{
		Name:  "Corp2",
		Email: "corp2@corp2.com",
		Plan:  models.Plan{ClientID: 2, Plan: "Enterprise"},
	},
}

func Load() {
	db := conn.GetDB()

	err := db.Debug().DropTableIfExists(&models.Client{}, &models.Plan{}).Error
	if err != nil {
		log.Fatal(err)
	}
	err = db.Debug().AutoMigrate(&models.Client{}, &models.Plan{}).Error
	if err != nil {
		log.Fatal(err)
	}

	for i := range clients {
		err = db.Debug().Model(&models.Client{}).Create(&clients[i]).Error
		if err != nil {
			log.Fatal(err)
		}
	}
}
