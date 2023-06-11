package main

import (
	"fmt"

	"github.com/Dimashey/comments-api/internal/comment"
	"github.com/Dimashey/comments-api/internal/db"
	transportHttp "github.com/Dimashey/comments-api/transport/http"
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
	httpHandler := transportHttp.NewHandler(commentService)

	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API Course")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
