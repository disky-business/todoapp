package controller

import (
	"net/http"

	"github.com/disky-business/todoapp/dto"
	"github.com/gin-gonic/gin"
)

func CreateTodoItem(c *gin.Context) {
	var todoDTO dto.TodoDTO
	if err := c.ShouldBindJSON(todoDTO); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Invalid request body",
			},
		)
	}

}
