package conn

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	opts := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", os.Getenv("HOST"),
		os.Getenv("DBPORT"), os.Getenv("DBUSER"), os.Getenv("DBNAME"), os.Getenv("PASSWORD"))

	conn, err := gorm.Open("postgres", opts)
	if err != nil {
		log.Printf("Failed to connect to DB.")
		os.Exit(100)
	}

	db = conn

}

func GetDB() *gorm.DB {
	return db
}
