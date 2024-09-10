package todo_test

import (
	"testing"

	"github.com/v420v/go-api/internal/todo/domain/models/todo"
)

func TestNewTodo(t *testing.T) {
	testCases := []struct {
		name          string
		title         string
		expectedError bool
	}{
		{
			name:          "Valid title",
			title:         "Test Todo",
			expectedError: false,
		},
		{
			name:          "Empty title",
			title:         "",
			expectedError: true,
		},
		{
			name:          "Title with 30 characters",
			title:         "This is a title with 30 chars.",
			expectedError: false,
		},
		{
			name:          "Title with more than 30 characters",
			title:         "This is a very long title that exceeds the 30 character limit",
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			todo, err := todo.NewTodo(tc.title)

			if tc.expectedError {
				if err == nil {
					t.Errorf("Expected an error, but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				if todo.Title() != tc.title {
					t.Errorf("Expected title %s, but got %s", tc.title, todo.Title())
				}

				if todo.Id() == "" {
					t.Error("Expected a non-empty ID")
				}

				if !todo.DeletedAt().IsZero() {
					t.Error("Expected DeletedAt to be zero")
				}
			}
		})
	}
}
