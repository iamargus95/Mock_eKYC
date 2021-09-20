package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	controllers "iamargus95/eKYC-service-gin/controllers"
)

func Connect() *gorm.DB {

	opts := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", os.Getenv("HOST"),
		os.Getenv("DBPORT"), os.Getenv("DBUSER"), os.Getenv("DBNAME"), os.Getenv("PASSWORD"))

	db, err := gorm.Open("postgres", opts)
	if err != nil {
		log.Printf("Failed to connect to DB.")
		os.Exit(100)
	}

	log.Printf("Connected to DB.")
	controllers.CreateClientTable(db)
	controllers.CreatePlansTable(db)
	controllers.InitDB(db)
	return db
}
