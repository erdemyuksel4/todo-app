package model

import "time"

type TodoStep struct {
	ID         int       `json:"id"`
	TodoListId int       `json:"listId" binding:"required"`
	Message    string    `json:"message"`
	IsDone     bool      `json:"is_done"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

type TodoList struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Percentage int       `json:"percentage"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}
