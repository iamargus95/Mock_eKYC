package main

import (
	postgres "iamargus95/eKYC-service-gin/postgres"
	routes "iamargus95/eKYC-service-gin/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	postgres.ClientConnect() //Client Database connected.
	postgres.PlansConnect()  //Plans Database connected.
	router := gin.Default()  // Init router
	routes.Routes(router)    // Route handlers & Endpoints
	log.Fatal(router.Run("localhost:8080"))
}
