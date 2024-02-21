package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/funukonta/todo-app/internal/model"
	"github.com/funukonta/todo-app/internal/service"
	"github.com/funukonta/todo-app/pkg"
)

type TodoHandler interface {
	CreateTask(w http.ResponseWriter, r *http.Request)
	GetTasks(w http.ResponseWriter, r *http.Request)
	UpdateTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
}

type todoHandler struct {
	service service.TodoService
}

func NewTodoHandler(service service.TodoService) TodoHandler {
	return &todoHandler{
		service: service,
	}
}

func (t *todoHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	taskReq := &model.Todo{}
	err := pkg.DecodeJsonReq(r, taskReq)
	if err != nil {
		pkg.JsonErr(w, http.StatusBadRequest, err)
	}

	taskDB, err := t.service.CreateTask(taskReq)
	if err != nil {
		pkg.JsonErr(w, http.StatusBadRequest, err)
	}

	pkg.JsonOK(w, http.StatusOK, "Berhasil buat task", taskDB)
}

func (t *todoHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	result, err := t.service.GetTasks()
	if err != nil {
		pkg.JsonErr(w, http.StatusBadRequest, err)
	}

	pkg.JsonOK(w, http.StatusOK, "Berhasil ambil data", result)
}

func (t *todoHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	req := &model.Todo{}
	err := pkg.DecodeJsonReq(r, req)
	if err != nil {
		pkg.JsonErr(w, http.StatusBadRequest, err)
	}

	result, err := t.service.UpdateTask(req)
	if err != nil {
		pkg.JsonErr(w, http.StatusBadRequest, err)
	}

	pkg.JsonOK(w, http.StatusOK, "Berhasil update", result)
}

func (t *todoHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	idStr := parts[len(parts)-1]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.JsonErr(w, http.StatusBadRequest, err)
	}

	err = t.service.DeleteTask(id)
	if err != nil {
		pkg.JsonErr(w, http.StatusBadRequest, err)
	}

	pkg.JsonOK(w, http.StatusBadRequest, "Berhasil delete", nil)
}
