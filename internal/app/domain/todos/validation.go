package todos

import (
	"errors"
	"strings"
)

func ValidateTodo(todo Todo) error {
	if todo.Title == "" || len(strings.TrimSpace(todo.Title)) == 0 {
		return errors.New("title is required")
	}
	if len(todo.Title) > 255 {
		return errors.New("title must be less than 255 characters")
	}
	return nil
}
