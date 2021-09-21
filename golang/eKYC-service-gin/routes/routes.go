package routes

import (
	ctrl "iamargus95/eKYC-service-gin/controllers/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/clients", ctrl.ClientsList)
	router.POST("/api/v1/signup", ctrl.PostClient) //--------- > Add SQL queries and seed data before implementing this
	router.NoRoute(notFound)
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
