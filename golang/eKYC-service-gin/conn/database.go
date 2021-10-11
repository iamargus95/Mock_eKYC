package conn

import (
	"fmt"
	"iamargus95/eKYC-service-gin/v1/models"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("HOST"), os.Getenv("DBPORT"), os.Getenv("DBUSER"),
		os.Getenv("DBNAME"), os.Getenv("PASSWORD"))

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB. ERROR: %v", err)
		os.Exit(100)
	}

	DB = conn

	sqlDB, _ := DB.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	DB.Debug().AutoMigrate(&models.Client{}, &models.Plan{},
		&models.SecretKey{}, &models.FileUpload{})

}
