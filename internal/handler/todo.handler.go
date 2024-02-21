package handler

import (
	"net/http"
	"strconv"

	"github.com/funukonta/todo-app/internal/model"
	"github.com/funukonta/todo-app/internal/service"
	"github.com/funukonta/todo-app/pkg"
	"github.com/gorilla/mux"
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
		return
	}

	taskDB, err := t.service.CreateTask(taskReq)
	if err != nil {
		pkg.JsonErr(w, http.StatusBadRequest, err)
		return
	}

	pkg.JsonOK(w, http.StatusOK, "Berhasil buat task", taskDB)
}

func (t *todoHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	result, err := t.service.GetTasks()
	if err != nil {
		pkg.WriteJson(w, http.StatusBadRequest, pkg.MsgError{Error: err})
		return
	}

	pkg.WriteJson(w, http.StatusOK, pkg.MsgOk{Data: result, Message: "Berhasil"})
}

func (t *todoHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	req := &model.Todo{}
	err := pkg.DecodeJsonReq(r, req)
	if err != nil {
		pkg.JsonErr(w, http.StatusBadRequest, err)
		return
	}

	result, err := t.service.UpdateTask(req)
	if err != nil {
		pkg.JsonErr(w, http.StatusBadRequest, err)
		return
	}

	pkg.JsonOK(w, http.StatusOK, "Berhasil update", result)
}

func (t *todoHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {

	idStr := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.JsonErr(w, http.StatusBadRequest, err)
		return
	}

	err = t.service.DeleteTask(id)
	if err != nil {
		pkg.JsonErr(w, http.StatusBadRequest, err)
		return
	}

	pkg.JsonOK(w, http.StatusBadRequest, "Berhasil delete", nil)
}
