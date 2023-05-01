package main

import (
	"fmt"
	"go-rest/internal/comment"
	"go-rest/internal/db"
	transportHttp "go-rest/internal/transport/http"

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

	// Services
	cmtService := comment.NewService(dbInstance)

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}
	// result, _ := cmtService.GetComment(context.Background(), "unqiue-id")
	// cmtService.PostComment(context.Background(), comment.Comment{
	// 	ID:     "test-comment-id",
	// 	Slug:   "manual-test",
	// 	Author: "Abhijit",
	// 	Body:   "hello beautiful world",
	// })

	return nil
}

func main() {

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
