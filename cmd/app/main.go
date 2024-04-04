package main

import (
	"context"
	"github.com/subosito/gotenv"
	"log"
	"todo/configs"
	_ "todo/docs"
	"todo/internal/controller/http"
	mongodb2 "todo/internal/repository/mongodb"
	"todo/internal/service"
	"todo/pkg/client/mongodb"
	"todo/pkg/web"
)

// @title todo-list app
// @version		1.0
// @description	Todo-list app
func main() {
	if err := run(); err != nil {
		log.Fatal(err.Error())
	}
}

func run() error {
	err := gotenv.Load(".env.local")
	if err != nil {
		return err
	}

	config, err := configs.New()
	if err != nil {
		return err
	}
	ctx := context.Background()

	db, err := mongodb.NewClient(ctx, *config)

	if err != nil {
		log.Println(err)
		panic("WTF")
	}

	repo := mongodb2.NewRepository(db)
	taskService := service.NewTaskService(repo)
	router := http.NewRouter(http.NewController(taskService))

	web.InitServer(router, *config)
	return nil
}
