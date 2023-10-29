package main

import (
	"log"

	todo "github.com/andres-website/todo-app/pkg"
	"github.com/andres-website/todo-app/pkg/handler"
	"github.com/andres-website/todo-app/pkg/repository"
	"github.com/andres-website/todo-app/pkg/service"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func main() {

	if err := initConfig(); err != nil {

		log.Fatalf("error initializing configs: %s ", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{

		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {

		log.Fatalf("failed initializing db: %s ", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {

		log.Fatalf("error accured while running http server: %s ", err.Error())
	}
}

func initConfig() error {

	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
