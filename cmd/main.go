package main

import (
	"log"
	"net/http"

	"github.com/funukonta/todo-app/internal/handler"
	"github.com/funukonta/todo-app/internal/repository"
	"github.com/funukonta/todo-app/internal/service"
	"github.com/funukonta/todo-app/pkg"
)

func main() {

	db, err := pkg.Connect()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.ServeMux{}

	repo := repository.NewPostgres(db)
	service := service.NewTodoService(repo)
	handler := handler.NewTodoHandler(service)

	mux.HandleFunc("POST /task", handler.CreateTask)

}
