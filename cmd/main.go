package main

import (
	"fmt"
	"os"

	todo "github.com/andres-website/todo-app/pkg"
	"github.com/andres-website/todo-app/pkg/handler"
	"github.com/andres-website/todo-app/pkg/repository"
	"github.com/andres-website/todo-app/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := initConfig(); err != nil {

		logrus.Fatalf("error initializing configs: %s ", err.Error())
	}

	if err := godotenv.Load(); err != nil {

		logrus.Fatalf("error loading env variavles: %s ", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{

		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {

		logrus.Fatalf("failed initializing db: %s ", err.Error())
	}

	err2 := db.Ping()

	if err2 != nil {
		fmt.Println("Ping не проходит main")
		return
	} else {

		fmt.Println("Ping проходит main")
	}

	repos := repository.NewRepository(db)

	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {

		logrus.Fatalf("error accured while running http server: %s ", err.Error())
	}
}

func initConfig() error {

	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
