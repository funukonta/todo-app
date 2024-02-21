package service

import (
	"github.com/funukonta/todo-app/internal/model"
	"github.com/funukonta/todo-app/internal/repository"
)

type TodoService interface {
	CreateTask(*model.Todo) (*model.Todo, error)
	GetTasks() ([]model.Todo, error)
	UpdateTask(taksUpdate *model.Todo) (*model.Todo, error)
	DeleteTask(id int) error
}

type todoService struct {
	repo repository.TodoRepo
}

func NewTodoService(repo repository.TodoRepo) TodoService {
	return &todoService{
		repo: repo,
	}
}

func (t *todoService) CreateTask(task *model.Todo) (*model.Todo, error) {
	task.Status = false
	return t.repo.CreateTask(task)

}

func (t *todoService) GetTasks() ([]model.Todo, error) {
	return t.repo.GetTasks()
}

func (t *todoService) UpdateTask(taksUpdate *model.Todo) (*model.Todo, error) {
	return t.repo.UpdateTask(taksUpdate)

}

func (t *todoService) DeleteTask(id int) error {
	return t.repo.DeleteTask(id)
}
