package enum

type TodoItemStatus int

const (
	DEFINED     TodoItemStatus = 1
	IN_PROGRESS TodoItemStatus = 2
	COMPLETED   TodoItemStatus = 3
)
