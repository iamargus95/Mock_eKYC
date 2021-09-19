package routes

import (
	"iamargus95/eKYC-service-gin/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/clients", controllers.GetClients)
	router.POST("/api/v1/signup", controllers.PostClient)
	router.GET("/client/:clientName", controllers.GetClientByName)
	router.PUT("/client/:clientName", controllers.EditPlanByName)
	router.DELETE("/client/:clientName", controllers.DeleteClient)
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
