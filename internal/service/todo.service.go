package service

import (
	"github.com/funukonta/todo-app/internal/model"
	"github.com/funukonta/todo-app/internal/repository"
)

type TodoService interface {
	CreateTask(*model.TODO) (*model.TODO, error)
}

type todoService struct {
	todoRepo repository.TodoRepo
}

func (t *todoService) CreateTask(task *model.TODO) (*model.TODO, error) {
	taskNew := model.TODO{
		TaskName: task.TaskName,
		DueDate:  task.DueDate,
		Priority: task.Priority,
		Status:   false,
	}

	taskInserted, err := t.todoRepo.CreateTask(&taskNew)
	if err != nil {
		return nil, err
	}

	return taskInserted, nil
}
