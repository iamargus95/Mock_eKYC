package routes

import (
	"log"
	"net/http"

	ctrl "iamargus95/eKYC-service-gin/v1/controllers"

	"github.com/gin-gonic/gin"
)

func SignupClient(r *gin.RouterGroup) {
	r.POST("/signup", ctrl.Signup)
}

func StartGin() {
	r := gin.Default() // Init router
	r.GET("/", welcome)
	r.NoRoute(notFound)
	SignupClient(r.Group("/api/v1"))
	log.Fatal(r.Run("localhost:8080"))
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
