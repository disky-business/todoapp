package service

import (
	"github.com/google/uuid"

	"github.com/disky-business/todoapp/dto"
	"github.com/disky-business/todoapp/models"
)



func CreateTodoItem(todoDTO dto.TodoDTO) {
	id := uuid.New()

	todoItem := models.TodoItem {
		ID: id,
		Title: todoDTO.Title,
		Description: todoDTO.Description,
		Priority: todoDTO.Priority,
		Status: todoDTO.Status,
		DueDate: todoDTO.DueDate,
	}
}
