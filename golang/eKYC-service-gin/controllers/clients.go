package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

//Struct of data
type AllData struct {
	Name  string `json:"name" validate:"required,min=2,max=50"`
	Email string `json:"email" validate:"required,email"`
	Tier  string `json:"tier"`
}

type Client struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Plans struct {
	Name string `json:"name"`
	Tier string `json:"tier"`
}

var dbClientConnect, dbPlanConnect *pg.DB

func InitializeDB(db *pg.DB) {
	dbClientConnect = db
	dbPlanConnect = db
}

func CreateClientTable(db *pg.DB) error {

	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	err := db.CreateTable(&Client{}, opts)
	if err != nil {
		log.Fatalf("Error while creating Clients table, ERROR : %v\n", err)
		return err
	}
	log.Printf("Clients table created.")
	return nil
}

func CreatePlansTable(db *pg.DB) error {

	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	err := db.CreateTable(&Plans{}, opts)
	if err != nil {
		log.Fatalf("Error while creating Plans table, ERROR : %v\n", err)
		return err
	}
	log.Printf("Plans table created.")
	return nil
}

//getClients responds with the list of all clients as JSON.
func GetClients(c *gin.Context) {

	var client []Client

	err := dbClientConnect.Model(&client).Select()

	if err != nil {
		log.Printf("Error while getting all clients, ERROR : %v.\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"Message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All clients.",
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
	tier := newData.Tier

	clientErr := dbClientConnect.Insert(&Client{
		Name:  name,
		Email: email,
	})

	planErr := dbPlanConnect.Insert(&Plans{
		Name: name,
		Tier: tier,
	})

	if clientErr != nil {
		log.Fatalf("Error while inserting new client into ClientDB, ERROR : %v.\n", clientErr)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Something went wrong.",
		})
		return
	}

	if planErr != nil {
		log.Fatalf("Error while inserting new plan into PlanDB, ERROR : %v.\n", planErr)
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

//Get data by name field
func GetClientByName(c *gin.Context) {

	name := c.Param("name")
	targetClient := &Client{Name: name}

	err := dbClientConnect.Select(targetClient)

	if err != nil {
		log.Printf("ERROR : %v.\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Client not found.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Target Client Found.",
		"data":    targetClient,
	})
}

func EditPlanByName(c *gin.Context) {

	clientName := c.Param("clientName")

	var targetPlan AllData

	c.BindJSON(&targetPlan)
	plan := targetPlan.Tier

	_, err := dbPlanConnect.Model(&Plans{}).Set("tier = ?", plan).Where("name = ?", clientName).Update()
	if err != nil {
		log.Printf("Error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Client plan changed successfully.",
	})
}

func DeleteClient(c *gin.Context) {

	clientName := c.Param("clientName")

	targetClient := &Client{Name: clientName}
	targetPlan := &Plans{Name: clientName}

	clientErr := dbClientConnect.Delete(targetClient)
	if clientErr != nil {
		log.Printf("Error while deleting the client from ClientDB, ERROR: %v.\n", clientErr)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong.",
		})
		return
	}

	planErr := dbPlanConnect.Delete(targetPlan)
	if planErr != nil {
		log.Printf("Error while deleting the client from PlanDB, ERROR: %v.\n", planErr)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Client was deleted successfully.",
	})
}
