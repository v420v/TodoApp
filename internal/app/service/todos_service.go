package service

import (
	"context"
	"errors"
	"time"

	"github.com/v420v/go-api/internal/app/domain/todos"
	"github.com/v420v/go-api/internal/app/repository"
)

func (s *Service) PostTodoService(todo todos.Todo) (todos.Todo, error) {
	if err := todos.ValidateTodo(todo); err != nil {
		return todos.Todo{}, errors.New("invalid todo")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return repository.InsertTodo(ctx, s.db, todo)
}

func (s *Service) TodoListService(page int, limit int) ([]todos.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	todoList, err := repository.SelectTodoList(ctx, s.db, page, limit)
	if err != nil {
		return nil, err
	}

	return todoList, nil
}

func (s *Service) DeleteTodo(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := repository.DeleteTodo(ctx, s.db, id)
	if err != nil {
		return err
	}

	return nil
}
