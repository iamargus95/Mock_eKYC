package conn

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		os.Getenv("HOST"), os.Getenv("DBPORT"), os.Getenv("DBUSER"),
		os.Getenv("DBNAME"), os.Getenv("PASSWORD"))

	conn, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB. ERROR: %v", err)
		os.Exit(100)
	}

	db = conn

}

func GetDB() *gorm.DB {
	return db
}
