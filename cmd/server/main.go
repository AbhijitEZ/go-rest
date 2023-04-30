package main

import (
	"fmt"
	"go-rest/internal/db"

	_ "github.com/joho/godotenv/autoload"
)

// Run - is going to responsible instantiation
func Run() error {
	fmt.Println("starting up application")

	dbInstance, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect the database")
		return err
	}

	fmt.Println("database connection successfull 🚀")

	if err := dbInstance.MigrateDB(); err != nil {
		fmt.Println("Failed to do the migrations")
		return err
	}

	return nil
}

func main() {

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}