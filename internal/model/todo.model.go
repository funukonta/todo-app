package model

import "time"

type TODO struct {
	Id        int       `json:"id" db:"id"`
	TaskName  string    `json:"taskname" db:"taskname"`
	DueDate   time.Time `json:"duedate" db:"duedate"`
	Priority  string    `json:"priority" db:"priority"`
	Status    bool      `json:"status" db:"status"`
	CreatedAt time.Time `json:"createdat" db:"createdat"`
	UpdatedAt time.Time `json:"updatedat" db:"updatedat"`
}
