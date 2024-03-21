package main

import (
	"context"
	"fmt"
	"main/pkg/handler"
	"main/pkg/repository"
	"main/pkg/service"
	"os"
	"os/signal"
	"syscall"
)

//mongodb+srv://lopatoandrey:Password@cluster0.efs83oj.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0
//mongodb+srv://lopatoandrey:<password>@cluster0.efs83oj.mongodb.net/

func main() {
	uri := "mongodb+srv://lopatoandrey:Password@cluster0.efs83oj.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	db, err := repository.NewMongoDb(uri)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if err := db.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	go func() {
		router := handlers.InitRoutes()
		router.Run("localhost:8080")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
