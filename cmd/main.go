package main

import (
	todo_app "github.com/quezzzzz/todo-app"
	"github.com/quezzzzz/todo-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo_app.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("eroor occured while ruuing http server: %s", err.Error())
	}

}
