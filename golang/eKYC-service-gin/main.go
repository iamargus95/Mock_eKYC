package main

import (
	routes "iamargus95/eKYC-service-gin/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() { //--------->Unfinished
	router := gin.Default() // Init router
	routes.Routes(router)   // Route handlers & Endpoints
	log.Fatal(router.Run("localhost:8080"))
}
