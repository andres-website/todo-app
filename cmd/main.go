package main

import (
	"log"

	todo "github.com/andres-website/todo-app/pkg"
	"github.com/andres-website/todo-app/pkg/handler"
	"github.com/andres-website/todo-app/pkg/repository"
	"github.com/andres-website/todo-app/pkg/service"
)

func main() {

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8008", handlers.InitRoutes()); err != nil {

		log.Fatalf("error accured while running http server: %s ", err.Error())
	}
}
