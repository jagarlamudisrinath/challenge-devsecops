package main

import (
	"challenge/api"
	"challenge/db"
	"challenge/dummy"
	"fmt"
	"os"
)

func init() {
	fmt.Println("DevSecOps challenge")

	if len(os.Getenv("POSTGRES_HOST")) > 0 {
		fmt.Println("Using postgresql DB driver")
		if err := db.PostgresConnection(); err != nil {
			fmt.Printf("Failed to connect to PostgreSQL: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println(os.Getenv("POSTGRES_HOST"))
		if err := db.SqliteConnector(); err != nil {
			fmt.Printf("Failed to connect to SQLite: %v\n", err)
			os.Exit(1)
		}
	}

	// Run migrations
	if err := db.Conn.AutoMigrate(&dummy.User{}); err != nil {
		fmt.Printf("Failed to run migrations: %v\n", err)
		os.Exit(2)
	}

	// Create the admin user
	var users []dummy.User
	db.Conn.Find(&users)
	if len(users) == 0 {
		fmt.Println("Could not find any users, bootstrapping an admin account")
		admin := dummy.User{
			Firstname: "Admin",
			Lastname:  "Istrator",
			Login:     "admin",
			Password:  "changeme",
		}
		if err := db.Conn.Create(&admin).Error; err != nil {
			fmt.Printf("Could not create admin user, reason %v", err)
			os.Exit(3)
		}
	} else {
		fmt.Println("Found users, skipping admin account bootstrapping")
	}

}

func main() {
	api.Start()
}
