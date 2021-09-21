package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

//Clients Structure.
type Clients struct {
	Model

	Name  string `gorm:"type:varchar(50)" json:"name" validate:"required"`
	Email string `gorm:"type:varchar(50)" json:"email" validate:"required,email"`
	Plans Plans
}

type Plans struct {
	Model

	ClientID uint
	Plan     string
}

//TableName return name of database table
func (d *Clients) TableName() string {
	return "clients"
}

func (d *Plans) TableName() string {
	return "plans"
}

func CreateClientsTable(db *gorm.DB) error {

	dbInfo := db.CreateTable(&Clients{})
	if dbInfo == nil {
		log.Fatalf("Error while creating Clients table.")
	}
	log.Printf("Clients table created.")
	return nil
}

func CreatePlansTable(db *gorm.DB) error {

	dbInfo := db.CreateTable(&Plans{})
	if dbInfo == nil {
		log.Fatalf("Error while creating Plans table.")
	}
	log.Printf("Plans table created.")
	return nil
}
