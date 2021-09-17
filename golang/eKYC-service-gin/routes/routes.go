package routes

import (
	"iamargus95/eKYC-service-gin/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/albums", controllers.GetAlbums)
	router.POST("/insert", controllers.PostAlbums)
	router.GET("/album/:albumID", controllers.GetAlbumById)
	router.PUT("/album/:albumID", controllers.EditAlbumByID)
	router.DELETE("/album/:albumID", controllers.DeleteAlbum)
	router.NoRoute(notFound)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome to the Draft API.",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  "404",
		"message": "Route not found.",
	})
	return
}
