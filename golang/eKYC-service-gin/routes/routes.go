package routes

import (
	ctrl "iamargus95/eKYC-service-gin/controllers/api/v1"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartGin() {
	router := gin.Default() // Init router
	router.GET("/", welcome)
	router.NoRoute(notFound)
	api := router.Group("api/v1")
	{
		api.GET("/clients", ctrl.ClientsList)
		// api.POST("/signup", ctrl.CreateClient)
	}
	log.Fatal(router.Run("localhost:8080"))
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome to the Draft API.",
	})
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  "404",
		"message": "Route not found.",
	})
}
