package todo

import "context"

type ITodoRepository interface {
	PostTodo(context.Context, Todo) error
	ListTodo(context.Context, int, int) ([]Todo, error)
	DeleteTodo(context.Context, string) error
}
