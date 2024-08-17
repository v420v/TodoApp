package services

import "github.com/v420v/go-api/internal/app/domain/todos"

type TodoServicer interface {
	PostTodoService(todos.Todo) (todos.Todo, error)
	TodoListService(page int, limit int) ([]todos.Todo, error)
	DeleteTodo(id int) error
}
