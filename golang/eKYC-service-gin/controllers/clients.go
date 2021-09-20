package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Struct of data
type AllData struct {
	Name  string
	Email string
	Plan  string
}

type Client struct {
	gorm.Model

	Name string `gorm:"name" validate:"required,min=2,max=50"`
}

type Plans struct {
	gorm.Model

	ClientID uint
	Client   Client
	Email    string `gorm:"email" validate:"required,email,unique_index"`
	Plan     string `gorm:"plan"`
}

var dbConnect *gorm.DB

func InitDB(db *gorm.DB) {
	dbConnect = db
}

func CreateClientTable(db *gorm.DB) error {

	dbInfo := db.CreateTable(&Client{})
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

//getClients responds with the list of all clients as JSON.
func GetClients(c *gin.Context) {

	var client Client
	dbInfo := dbConnect.Raw("SELECT * FROM clients").Scan(&client)
	if dbInfo == nil {
		log.Printf("Error while getting all clients, ERROR : %v.\n", dbInfo)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"Message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All clients",
		"data":    client,
	})
}

//Add client entry to the Database
func PostClient(c *gin.Context) {

	var newData AllData
	//call Bind JSON to bind the received JSON to newClient
	if err := c.BindJSON(&newData); err != nil {
		return
	}

	name := newData.Name
	email := newData.Email
	plan := newData.Plan

	planErr := dbConnect.Create(&Plans{
		Client: Client{Name: name},
		Email:  email,
		Plan:   plan,
	})

	if planErr == nil {
		log.Fatalf("Error while inserting new plan into Plan Table, ERROR : %v.\n", planErr)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Something went wrong.",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":    http.StatusOK,
		"accessKey": "10-Char-String",
		"secretKey": "20-Char-String",
	})
}
