package main

import (
	"context"
	"fmt"

	"github.com/Dimashey/comments-api/internal/comment"
	"github.com/Dimashey/comments-api/internal/db"
)

// Run - is going to be responsible
// the instantiation and startup of our
// go application
func Run() error {
	fmt.Println("starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")

		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failded to migrate datbase")
		return err
	}

	commentService := comment.NewService(db)

	fmt.Println(commentService.GetComment(context.Background(), "e0e2ae8f-4af4-44c7-a7fa-798d6dc6e394"))
	return nil
}

func main() {
	fmt.Println("Go REST API Course")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
