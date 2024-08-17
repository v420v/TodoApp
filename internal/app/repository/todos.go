package repository

import (
	"context"
	"time"

	"github.com/uptrace/bun"
	"github.com/v420v/go-api/internal/app/domain/todos"
)

func InsertTodo(ctx context.Context, db *bun.DB, todo todos.Todo) (todos.Todo, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return todos.Todo{}, err
	}

	now := time.Now().In(jst)
	todo.CreatedAt = now
	todo.UpdatedAt = now

	_, err = db.NewInsert().Model(&todo).Returning("*").Exec(ctx)
	if err != nil {
		return todos.Todo{}, err
	}
	return todo, nil
}

func SelectTodoList(ctx context.Context, db *bun.DB, page int, limit int) ([]todos.Todo, error) {
	var todos = []todos.Todo{}

	err := db.NewSelect().
		Model(&todos).
		Where("deleted_at IS NULL").
		Order("created_at DESC").
		Offset((page - 1) * limit).
		Limit(limit).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return todos, nil
}

func DeleteTodo(ctx context.Context, db *bun.DB, id int) error {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}

	now := time.Now().In(jst)

	_, err = db.NewUpdate().Model(&todos.Todo{}).Set("deleted_at = ?", now).Where("todo_id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
