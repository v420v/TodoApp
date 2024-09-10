package todo

import (
	"errors"
	"time"
	"unicode/utf8"

	"github.com/bwmarrin/snowflake"
)

type Todo struct {
	id        string
	title     string
	createdAt time.Time
	updatedAt time.Time
	deletedAt time.Time
}

func NewId() (string, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return "", err
	}

	id := node.Generate().Base64()

	return id, nil
}

func NewTodo(title string) (Todo, error) {
	if utf8.RuneCountInString(title) > 30 {
		return Todo{}, errors.New("Todo title must be under 30 characters")
	}
	if utf8.RuneCountInString(title) == 0 {
		return Todo{}, errors.New("Todo title must have at least one character")
	}

	id, err := NewId()
	if err != nil {
		return Todo{}, err
	}

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return Todo{}, err
	}

	now := time.Now().In(jst)

	return Todo{
		id:        id,
		title:     title,
		createdAt: now,
		updatedAt: now,
	}, nil
}

func ConvertDaoTodoToTodo(id string, title string, createdAt time.Time, updatedAt time.Time, deletedAt time.Time) Todo {
	return Todo{
		id:        id,
		title:     title,
		createdAt: createdAt,
		updatedAt: updatedAt,
		deletedAt: deletedAt,
	}
}

func (t Todo) Id() string {
	return t.id
}

func (t Todo) Title() string {
	return t.title
}

func (t Todo) CreatedAt() time.Time {
	return t.createdAt
}

func (t Todo) UpdatedAt() time.Time {
	return t.updatedAt
}

func (t Todo) DeletedAt() time.Time {
	return t.deletedAt
}
