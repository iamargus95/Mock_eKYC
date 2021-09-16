package main

//Tutorial available at https://golang.org/doc/tutorial/web-service-gin
import (
	"iamargus95/eKYC-service-gin/controllers"
	"iamargus95/eKYC-service-gin/postgres"

	"github.com/gin-gonic/gin"
)

func main() {
	postgres.Connect() //Database connected.
	router := gin.Default()
	router.GET("/albums", controllers.GetAlbums)
	router.GET("/albums/:id", controllers.GetAlbumById)
	router.POST("/albums", controllers.PostAlbums)

	router.Run("localhost:8080")
}
