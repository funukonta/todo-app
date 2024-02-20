package repository

import (
	"context"

	"github.com/funukonta/todo-app/internal/model"
	"github.com/jmoiron/sqlx"
)

type TodoRepo interface {
	// SelectTask()
	// SelectTasks()
	CreateTask(*model.TODO) (*model.TODO, error)
	// UpdateTask()
	// DeleteTask()
}

type todoRepo struct {
	db *sqlx.DB
}

func NewPostgres(db *sqlx.DB) TodoRepo {
	return &todoRepo{
		db: db,
	}
}

func (t *todoRepo) CreateTask(task *model.TODO) (*model.TODO, error) {
	query := `INSERT into taks (taskname,duedate,priority,status) values 
	($1,$2,$3,$4)
	RETURNING id,taskname,duedate,priority,status,createdat
	`

	tx, err := t.db.BeginTxx(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	createdTask := new(model.TODO)
	err = tx.QueryRowx(query, task.TaskName, task.DueDate, task.Priority, task.Status).StructScan(createdTask)
	if err != nil {
		return nil, err
	}

	return createdTask, nil
}
