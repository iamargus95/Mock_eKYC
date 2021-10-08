package routes

import (
	"log"
	"net/http"

	authtoken "iamargus95/eKYC-service-gin/jwt"
	"iamargus95/eKYC-service-gin/middlewares"
	ctrl "iamargus95/eKYC-service-gin/v1/controllers"

	"github.com/gin-gonic/gin"
)

func SignupClient(r *gin.RouterGroup) {
	r.POST("/signup", ctrl.Signup)
}

func ImageUpload(r *gin.RouterGroup) {
	r.POST("/image", ctrl.Image, middlewares.EnsureLoggedIn(authtoken.JWTService()))
}

func StartGin() {
	r := gin.Default() // Init router
	r.GET("/", Welcome)
	r.NoRoute(NotFound)
	routerGroup := r.Group("/api/v1")
	SignupClient(routerGroup)
	ImageUpload(routerGroup)
	log.Fatal(r.Run("localhost:8080"))
}

func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Draft API.",
	})
}

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Route not found.",
	})
}
