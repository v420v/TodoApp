package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/v420v/go-api/internal/todo/app/command"
	"github.com/v420v/go-api/internal/todo/app/service"
	"github.com/v420v/go-api/internal/todo/domain/models/todo"
)

type TodoController struct {
	todoService *service.TodoApplicationService
}

func NewTodoController(todoService *service.TodoApplicationService) *TodoController {
	return &TodoController{
		todoService: todoService,
	}
}

type TodoResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func toTodoResponse(t todo.Todo) TodoResponse {
	return TodoResponse{
		ID:        t.Id(),
		Title:     t.Title(),
		CreatedAt: t.CreatedAt(),
		UpdatedAt: t.UpdatedAt(),
		DeletedAt: t.DeletedAt(),
	}
}

// @Summary	Get Todo list
// @Description	get a list of TODOs that are not deleted.
// @Tags			todo
// @Accept			json
// @Produce			json
// @Param			page query int false "page number"
// @Param			limit query	int false "number of items per page"
// @Success			200 {array} todo.Todo
// @Router			/todos [get]
func (c *TodoController) PostTodo(w http.ResponseWriter, req *http.Request) {
	type todoInput struct {
		Title string `json:"title"`
	}

	input := todoInput{}

	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(&input)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	command := command.TodoCreateCommand{
		Title: input.Title,
	}

	err = c.todoService.Create(command)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary	Get Todo list
// @Description	get a list of TODOs that are not deleted.
// @Tags			todo
// @Accept			json
// @Produce			json
// @Param			page query int false "page number"
// @Param			limit query	int false "number of items per page"
// @Success			200 {array} todo.Todo
// @Router			/todos [get]
func (c *TodoController) GetList(w http.ResponseWriter, req *http.Request) {
	pageStr := req.URL.Query().Get("page")
	limitStr := req.URL.Query().Get("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	command := command.TodoGetListCommand{
		Page:  page,
		Limit: limit,
	}

	todoList, err := c.todoService.GetList(command)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	responseTodoList := []TodoResponse{}
	for _, t := range todoList {
		responseTodoList = append(responseTodoList, toTodoResponse(t))
	}

	json.NewEncoder(w).Encode(responseTodoList)
	w.WriteHeader(http.StatusOK)
}

// @Summary Delete todo
// @Description	delete todo by id
// @Tags			todo
// @Accept			json
// @Param			id	path int true "user id"
// @Success			200
// @Router			/todos/{id}/delete [post]
func (c *TodoController) DeleteTodo(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	command := command.TodoDeleteCommand{
		Id: id,
	}

	if err := c.todoService.Delete(command); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
