package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

//Struct of data
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func CreateAlbumTable(db *pg.DB) error {

	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	err := db.CreateTable(&album{}, opts)
	if err != nil {
		log.Fatalf("Error while creating album table, ERROR : %v\n", err)
		return err
	}
	log.Printf("Album table created.")
	return nil
}

var dbConnect *pg.DB

func InitializeDB(db *pg.DB) {

	dbConnect = db
}

//getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {

	var albums []album

	err := dbConnect.Model(&albums).Select()

	if err != nil {
		log.Printf("Error while getting all albums, ERROR : %v.\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"Message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All albums.",
		"data":    albums,
	})
}

//Add an entry to the Database
func PostAlbums(c *gin.Context) {

	var newAlbum album
	//call Bind JSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	id := newAlbum.ID
	title := newAlbum.Title
	artist := newAlbum.Artist
	price := newAlbum.Price

	inserterr := dbConnect.Insert(&album{
		ID:     id,
		Title:  title,
		Artist: artist,
		Price:  price,
	})

	if inserterr != nil {
		log.Fatalf("Error while inserting new album into db, ERROR : %v.\n", inserterr)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong.",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "New album created successfully.",
	})
}

//Get data by id field
func GetAlbumById(c *gin.Context) {

	id := c.Param("albumID")
	targetAlbum := &album{ID: id}

	err := dbConnect.Select(targetAlbum)

	if err != nil {
		log.Printf("Error while getting a single album, ERROR : %v.\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Album not found.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Target Album",
		"data":    targetAlbum,
	})
}

func EditAlbumByID(c *gin.Context) {

	albumId := c.Param("albumID")

	var targetAlbum album

	c.BindJSON(&targetAlbum)
	price := targetAlbum.Price

	_, err := dbConnect.Model(&album{}).Set("price = ?", price).Where("id = ?", albumId).Update()
	if err != nil {
		log.Printf("Error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Album Price Edited Successfully",
	})
}

func DeleteAlbum(c *gin.Context) {
	albumId := c.Param("albumID")
	targetAlbum := &album{ID: albumId}

	err := dbConnect.Delete(targetAlbum)
	if err != nil {
		log.Printf("Error while deleting a single album, ERROR: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Album deleted successfully",
	})
}
