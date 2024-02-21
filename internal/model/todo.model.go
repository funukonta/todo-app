package model

import "time"

type Todo struct {
	Id        int       `json:"id" db:"id"`
	TaskName  string    `json:"taskname" db:"taskname"`
	Desc      string    `json:"desc" db:"desc"`
	DueDate   time.Time `json:"duedate" db:"duedate"`
	Priority  int       `json:"priority" db:"priority"`
	Status    bool      `json:"status" db:"status"`
	CreatedAt time.Time `json:"createdat,omitempty" db:"createdat"`
	UpdatedAt time.Time `json:"updatedat,omitempty" db:"updatedat"`
}
