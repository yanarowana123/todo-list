package main

import (
	"context"
	"github.com/subosito/gotenv"
	"github.com/yanarowana123/todo-list/configs"
	_ "github.com/yanarowana123/todo-list/docs"
	"github.com/yanarowana123/todo-list/internal/controller/http"
	mongodb2 "github.com/yanarowana123/todo-list/internal/repository/mongodb"
	"github.com/yanarowana123/todo-list/internal/service"
	"github.com/yanarowana123/todo-list/pkg/client/mongodb"
	"github.com/yanarowana123/todo-list/pkg/web"
	"log"
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
		return err
	}
	err = mongodb.CreateIndexes(ctx, db)
	if err != nil {
		return err
	}

	repo := mongodb2.NewRepository(db)
	taskService := service.NewTaskService(repo)
	router := http.NewRouter(http.NewController(taskService))

	web.InitServer(router, *config)
	return nil
}
