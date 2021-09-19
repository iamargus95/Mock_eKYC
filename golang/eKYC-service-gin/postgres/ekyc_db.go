package postgres

import (
	"log"
	"os"

	"github.com/go-pg/pg"

	controllers "iamargus95/eKYC-service-gin/controllers"
)

func ClientConnect() *pg.DB {
	opts := &pg.Options{
		User:     os.Getenv("DBUSER"),
		Password: os.Getenv("PASSWORD"),
		Addr:     os.Getenv("DBPORT"),
		Database: os.Getenv("DBCLIENT"),
	}

	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		log.Printf("Failed to connect to ClientDB.")
		os.Exit(100)
	}

	log.Printf("Connected to ClientDB.")
	controllers.CreateClientTable(db) // ---------> Fails here
	controllers.InitializeDB(db)
	return db
}

func PlansConnect() *pg.DB {
	opts := &pg.Options{
		User:     os.Getenv("DBUSER"),
		Password: os.Getenv("PASSWORD"),
		Addr:     os.Getenv("DBPORT"),
		Database: os.Getenv("DBPLAN"),
	}

	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		log.Printf("Failed to connect to PlanDB.")
		os.Exit(100)
	}

	log.Printf("Connected to PlanDB.")
	controllers.CreatePlansTable(db)
	controllers.InitializeDB(db)
	return db
}
