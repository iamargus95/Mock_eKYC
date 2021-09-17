package postgres

import (
	"log"
	"os"

	"github.com/go-pg/pg"

	controllers "iamargus95/eKYC-service-gin/controllers"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     os.Getenv("DBUSER"),
		Password: os.Getenv("PASSWORD"),
		Addr:     os.Getenv("DBPORT"),
		Database: os.Getenv("DBNAME"),
	}

	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		log.Printf("Failed to connect to DB.")
		os.Exit(100)
	}

	log.Printf("Connected to DB.")
	controllers.CreateAlbumTable(db)
	controllers.InitializeDB(db)
	return db
}
