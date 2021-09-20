package main

import (
	postgres "iamargus95/eKYC-service-gin/postgres"
	routes "iamargus95/eKYC-service-gin/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db := postgres.Connect() //Client Database connected.
	defer db.Close()

	router := gin.Default() // Init router
	routes.Routes(router)   // Route handlers & Endpoints
	log.Fatal(router.Run("localhost:8080"))
}
