package dto

import (
	"github.com/disky-business/todoapp/enum"
)

type TodoDTO struct {
	Title        string                `json:"title"`
	Description  string                `json:"description"`
	Priority     enum.TodoItemPriority `json:"priority"`
	Status       enum.TodoItemStatus   `json:"status"`
	DueDate      string                `json:"due_date"`
}
