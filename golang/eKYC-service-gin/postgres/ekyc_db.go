package postgres

import (
	"log"
	"os"

	"github.com/go-pg/pg"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		Addr:     os.Getenv("DBPORT"),
		Database: os.Getenv("NAME"),
	}

	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		log.Printf("Failed to connect to DB.")
		os.Exit(100)
	}

	log.Printf("Connected to DB.")

	return db
}
