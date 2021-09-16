package main

//Tutorial available at https://golang.org/doc/tutorial/web-service-gin
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Struct of data

type album struct {
	ID     string  `json"id"`
	Title  string  `json"title"`
	Artist string  `json"artist"`
	Price  float64 `json"price"`
}

//albums slice to seed record albm data

var albums = []album{
	{"1", "Blue Train", "John Coltrane", 56.99},
	{"2", "Jeru", "Gerry Mulligan", 17.99},
	{"3", "Sarah Vaughan and Clifford Brown", "Sarah Vaughan", 39.99},
}

//getAlbums responds with the list of all albums as JSON.

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	//call Bind JSON to bind the received JSON to newAlbum

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
