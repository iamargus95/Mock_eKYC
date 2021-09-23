package routes

import (
	"log"
	"net/http"

	ctrl "iamargus95/eKYC-service-gin/controllers/api/v1"
	v1s "iamargus95/eKYC-service-gin/services/api/v1"

	"github.com/gin-gonic/gin"
)

var (
	signupService    v1s.SignupService     = v1s.New()
	signupController ctrl.SignupController = ctrl.New(signupService)
)

func StartGin() {
	router := gin.Default() // Init router
	router.GET("/", welcome)
	router.NoRoute(notFound)
	api := router.Group("api/v1")
	{
		api.POST("/signup", func(ctx *gin.Context) {
			ctx.JSON(200, signupController.Save(ctx))
		})
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
