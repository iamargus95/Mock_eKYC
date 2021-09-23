package infra

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/joho/godotenv"
)

var m *migrate.Migrate

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	Dsn := fmt.Sprintf("postgres://postgres:%s@localhost:%s/%s?sslmode=disable", os.Getenv("PASSWORD"), os.Getenv("DBPORT"), os.Getenv("DBNAME"))
	m, err = migrate.New("file:///home/suraj/git/Mock_eKYC/golang/eKYC-service-gin/schema/migrations/sql", Dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}

func GetMigrationTool() *migrate.Migrate {
	return m
}
