package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"os"
)

var Conn *gorm.DB
var err error

var DBHost string
var DBUser string

func init() {

	if len(os.Getenv("POSTGRES_USER")) != 0 {
		DBUser = os.Getenv("POSTGRES_USER")
	} else {
		DBUser = "challenge"
	}

	if len(os.Getenv("POSTGRES_HOST")) != 0 {
		DBHost = os.Getenv("POSTGRES_HOST")
	} else {
		DBHost = "localhost"
	}

}

func SqliteConnector() error {

	Conn, err = gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func PostgresConnection() error {

	dsn := fmt.Sprintf("host=%v user=%v password=6199178B-28C5-4960-89FD-B1E55E0044E6 dbname=challenge port=5432 sslmode=disable", DBHost, DBUser)
	Conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	return nil

}
