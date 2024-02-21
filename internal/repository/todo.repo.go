package repository

import (
	"context"
	"fmt"

	"github.com/funukonta/todo-app/internal/model"
	"github.com/jmoiron/sqlx"
)

type TodoRepo interface {
	CreateTask(*model.Todo) (*model.Todo, error)
	GetTasks() ([]model.Todo, error)
	UpdateTask(taksUpdate *model.Todo) (*model.Todo, error)
	DeleteTask(id int) error
}

type todoRepo struct {
	db *sqlx.DB
}

func NewRepoPostgres(db *sqlx.DB) TodoRepo {
	return &todoRepo{
		db: db,
	}
}

func (t *todoRepo) CreateTask(task *model.Todo) (*model.Todo, error) {
	query := `INSERT into tasks (taskname,description,duedate,priority,status) values ($1,$2,$3,$4,$5)
	RETURNING id,taskname,description,duedate,priority,status,createdat`

	tx, err := t.db.BeginTxx(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	output := &model.Todo{}
	err = tx.QueryRowx(query,
		task.TaskName,
		task.Desc,
		task.DueDate,
		task.Priority,
		task.Status,
	).StructScan(output)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (t *todoRepo) GetTasks() ([]model.Todo, error) {
	query := `select * from tasks order by duedate desc;`

	rows, err := t.db.Queryx(query)
	if err != nil {
		return nil, err
	}

	results := []model.Todo{}
	for rows.Next() {
		resultRow := model.Todo{}
		err := rows.StructScan(&resultRow)
		if err != nil {
			return nil, err
		}

		results = append(results, resultRow)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("data not found")
	}

	return results, nil
}

func (t *todoRepo) UpdateTask(taksUpdate *model.Todo) (*model.Todo, error) {
	query := `UPDATE tasks set taskname=$1,description=$2,duedate=$3,priority=$4,status=$5,updatedat=now() where id=$6
	returning RETURNING id,taskname,description,duedate,priority,status,updatedat`

	tx, err := t.db.BeginTxx(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	updatedTask := &model.Todo{}
	err = tx.QueryRowx(query,
		taksUpdate.TaskName,
		taksUpdate.Desc,
		taksUpdate.DueDate,
		taksUpdate.Priority,
		taksUpdate.Status,
		taksUpdate.Id,
	).Scan(updatedTask)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (t *todoRepo) DeleteTask(id int) error {
	query := `delete from tasks where id=$1`

	tx, err := t.db.BeginTxx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(query, id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
