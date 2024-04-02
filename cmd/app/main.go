package main

import (
	"context"
	"log"
	"todo/internal/controller/http"
	mongodb2 "todo/internal/repository/mongodb"
	"todo/internal/service"
	"todo/pkg/client/mongodb"
	"todo/pkg/web"
)

func main() {
	//gotenv.Load()
	if err := run(); err != nil {
		log.Fatal(err.Error())
	}
}

func run() error {

	ctx := context.Background()
	db, err := mongodb.NewClient(ctx, "localhost", "27018", "", "", "todo-list", "")

	if err != nil {
		panic("WTF")
	}

	repo := mongodb2.NewRepository(db)
	//repo := inmemory.NewRepository()
	service := service.NewTaskService(repo)
	router := http.NewRouter(http.NewController(service))

	web.InitServer(router)
	return nil

	//config, err := configs.New()
	//if err != nil {
	//	return err
	//}
	//
	//validate := validator.New()
	//
	//repositoryManager := repositories.NewManager(*config)
	//
	//serviceManager := services.NewManager(*repositoryManager, *config)
	//
	//handlerManager := handler.NewManager(*serviceManager, validate)
	//
	//r := mux.NewRouter()
	//router := web.InitRouter(r, *handlerManager)
	//
	//web.InitServer(*config, router)
	//return nil
}
