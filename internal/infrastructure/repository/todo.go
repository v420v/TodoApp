package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/uptrace/bun"
	"github.com/v420v/go-api/internal/infrastructure/dao"
	"github.com/v420v/go-api/internal/todo/domain/models/todo"
)

type TodoRepository struct {
	db *bun.DB
}

func NewTodoRepository(db *bun.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (r *TodoRepository) PostTodo(ctx context.Context, todoData todo.Todo) error {
	t := dao.Todo{
		Id:        todoData.Id(),
		Title:     todoData.Title(),
		CreatedAt: todoData.CreatedAt(),
		UpdatedAt: todoData.UpdatedAt(),
		DeletedAt: todoData.DeletedAt(),
	}

	_, err := r.db.NewInsert().Model(&t).Returning("*").Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (r *TodoRepository) ListTodo(ctx context.Context, page int, limit int) ([]todo.Todo, error) {
	todoList := []dao.Todo{}

	err := r.db.NewSelect().
		Model(&todoList).
		Where("deleted_at IS NULL").
		Order("created_at DESC").
		Offset((page - 1) * limit).
		Limit(limit).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	resultTodoList := []todo.Todo{}
	for _, todoDao := range todoList {
		resultTodoList = append(resultTodoList, todo.ConvertDaoTodoToTodo(
			todoDao.Id,
			todoDao.Title,
			todoDao.CreatedAt,
			todoDao.UpdatedAt,
			todoDao.DeletedAt,
		))
	}

	fmt.Println(resultTodoList)

	return resultTodoList, nil
}

func (r *TodoRepository) DeleteTodo(ctx context.Context, id string) error {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}

	now := time.Now().In(jst)

	_, err = r.db.NewUpdate().Model(&todo.Todo{}).Set("deleted_at = ?", now).Where("todo_id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
