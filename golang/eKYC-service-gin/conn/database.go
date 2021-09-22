package conn

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var db *sql.DB

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	Dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		os.Getenv("HOST"), os.Getenv("DBPORT"), os.Getenv("DBUSER"),
		os.Getenv("DBNAME"), os.Getenv("PASSWORD"))

	conn, err := sql.Open("postgres", Dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB. ERROR: %v", err)
		os.Exit(100)
	}

	db = conn

}

func GetDB() *sql.DB {
	return db
}
