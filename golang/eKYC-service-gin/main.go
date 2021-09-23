package main

import (
	"iamargus95/eKYC-service-gin/routes"
	seed "iamargus95/eKYC-service-gin/seeder"
)

func main() {
	seed.Load()
	routes.StartGin()
}
