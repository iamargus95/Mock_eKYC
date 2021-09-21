package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB

//Model is sample of common table structure
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at" sql:DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at" sql:DEFAULT:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

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

	db.Debug().AutoMigrate(
		&Clients{},
		&Plans{},
	)
}

func GetDB() *gorm.DB {
	return db
}
