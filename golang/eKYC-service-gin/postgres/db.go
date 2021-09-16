package postgres

import (
	"log"
	"os"

	"github.com/go-pg/pg"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "db_username",
		Password: "db_password",
		Addr:     "localhost:5452",
		Database: "db_dbname",
	}

	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		log.Printf("Failed to connect to DB.")
		os.Exit(100)
	}

	log.Printf("Connected to DB.")

	return db
}
