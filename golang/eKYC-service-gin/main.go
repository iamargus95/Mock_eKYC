package main

//Tutorial available at https://golang.org/doc/tutorial/web-service-gin
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Struct of data

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

//albums slice to seed record albm data

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

//getAlbums responds with the list of all albums as JSON.

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

//Add an entry to the Database
func postAlbums(c *gin.Context) {
	var newAlbum album

	//call Bind JSON to bind the received JSON to newAlbum

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

//Get data by id field

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	//Loop over the list of albums, looking for an album whose ID values match the parameter.

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found."})
}
