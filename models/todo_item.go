package models

import (
	"github.com/google/uuid"

	"github.com/disky-business/todoapp/enum"
)

type TodoItem struct {
	ID 		     uuid.UUID             `json:"id"`
	Title        string                `json:"title"`
	Description  string                `json:"description"`
	Priority     enum.TodoItemPriority `json:"priority"`
	Status       enum.TodoItemStatus   `json:"status"`
	DueDate      string                `json:"due_date"`
}
