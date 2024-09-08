package service

import (
	"context"
	"time"

	"github.com/v420v/go-api/internal/todo/app/command"
	"github.com/v420v/go-api/internal/todo/domain/models/todo"
)

type TodoApplicationService struct {
	todoRepository todo.ITodoRepository
}

func NewTodoApplicationService(todoRepository todo.ITodoRepository) *TodoApplicationService {
	return &TodoApplicationService{
		todoRepository: todoRepository,
	}
}

func (s *TodoApplicationService) GetList(cmd command.TodoGetListCommand) ([]todo.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	todoList, err := s.todoRepository.ListTodo(ctx, cmd.Page, cmd.Limit)

	if err != nil {
		return []todo.Todo{}, err
	}

	return todoList, nil
}

func (s *TodoApplicationService) Create(cmd command.TodoCreateCommand) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	todoData, err := todo.NewTodo(cmd.Title)
	if err != nil {
		return err
	}

	err = s.todoRepository.PostTodo(ctx, todoData)

	if err != nil {
		return err
	}

	return nil
}

func (s *TodoApplicationService) Delete(cmd command.TodoDeleteCommand) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.todoRepository.DeleteTodo(ctx, cmd.Id)

	if err != nil {
		return err
	}

	return nil
}
