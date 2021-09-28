package routes

import (
	"log"
	"net/http"

	ctrl "iamargus95/eKYC-service-gin/v1/controllers"

	"github.com/gin-gonic/gin"
)

func StartGin() {
	router := gin.Default() // Init router
	router.GET("/", welcome)
	router.NoRoute(notFound)
	api := router.Group("api/v1")
	{
		api.POST("/signup", ctrl.Signup)
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
