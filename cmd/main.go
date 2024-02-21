package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/funukonta/todo-app/internal/handler"
	"github.com/funukonta/todo-app/internal/repository"
	"github.com/funukonta/todo-app/internal/service"
	"github.com/funukonta/todo-app/pkg"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load Env
	evnPath := filepath.Join("..", ".env")
	err := godotenv.Load(evnPath)
	if err != nil {
		log.Fatal("Error load .env" + err.Error())
	}

	db := pkg.NewPostgreDB()
	if err != nil {
		log.Println(err)
	}

	repo := repository.NewRepoPostgres(db)
	service := service.NewTodoService(repo)
	handler := handler.NewTodoHandler(service)

	mux := mux.NewRouter()

	mux.HandleFunc("/todo", handler.CreateTask).Methods("POST")
	mux.HandleFunc("/todo", handler.GetTasks).Methods("GET")
	mux.HandleFunc("/todo", handler.UpdateTask).Methods("PUT")
	mux.HandleFunc("/todo", handler.DeleteTask).Methods("DELETE")

	port := os.Getenv("SERVER_PORT")
	log.Println("Server Start at :", port)
	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}

}
