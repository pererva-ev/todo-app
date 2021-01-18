package main

import (
	"log"

	"github.com/pererva-ev/todo-app"
	"github.com/pererva-ev/todo-app/pkg/handler"
)

func main() {
	handlers := handler.Handler{}
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured whilee running http server: %s", err.Error())
	}
}
