package main

import (
	"iamargus95/eKYC-service-gin/conn"
	routes "iamargus95/eKYC-service-gin/v1/routes"
)

func main() {
	conn.SetupDB()
	routes.StartGin()
}
