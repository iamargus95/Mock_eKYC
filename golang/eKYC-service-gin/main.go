package main

import (
	"iamargus95/eKYC-service-gin/conn"
	"iamargus95/eKYC-service-gin/v1/models"
	v1routes "iamargus95/eKYC-service-gin/v1/routes"
	"time"
)

func main() {

	db := conn.GetDB()
	defer db.Close()

	sqlDB := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	db.Debug().AutoMigrate(&models.Client{}, &models.Plan{}, &models.MiddleMap{})
	v1routes.StartGin()
}
