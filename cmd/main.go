package main

import (
	"log"

	todo "github.com/andres-website/todo-app/pkg"
	"github.com/andres-website/todo-app/pkg/handler"
)

func main() {

	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8008", handlers.InitRoutes()); err != nil {

		log.Fatalf("error accured while running http server: %s ", err.Error())
	}
}
