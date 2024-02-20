package handler

import (
	"net/http"

	"github.com/funukonta/todo-app/internal/model"
	"github.com/funukonta/todo-app/internal/service"
	"github.com/funukonta/todo-app/pkg"
)

type TodoHandler interface {
	CreateTask(http.ResponseWriter, *http.Request)
}

type todoHandler struct {
	todoService service.TodoService
}

func (t *todoHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	taskNew := model.TODO{}
	err := pkg.DecodeJsonReq(r, &taskNew)
	if err != nil {
		pkg.WriteErrorJson(w, http.StatusBadRequest, err)
	}

	taskInserted, err := t.todoService.CreateTask(&taskNew)
	if err != nil {
		pkg.WriteErrorJson(w, http.StatusExpectationFailed, err)
	}

	pkg.WriteSuccessJson(w, http.StatusOK, taskInserted, "Berhasil Insert")
}
