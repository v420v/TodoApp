package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/v420v/go-api/internal/app/controller/services"
	"github.com/v420v/go-api/internal/app/domain/todos"
)

type TodoController struct {
	service services.TodoServicer
}

func NewTodoController(service services.TodoServicer) *TodoController {
	return &TodoController{service: service}
}

//	@Summary		Post todo
//	@Description	insert todo to db
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			todo	body		todos.Todo	true	"Todo object"
//	@Success		200		{object}	todos.Todo
//	@Router			/todos [post]
func (c *TodoController) PostTodoHandler(w http.ResponseWriter, req *http.Request) {
	var todo todos.Todo
	decoder := json.NewDecoder(req.Body)

	if err := decoder.Decode(&todo); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	todo, err := c.service.PostTodoService(todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(todo)
	w.WriteHeader(http.StatusOK)
}

//	@Summary		Get Todo list
//	@Description	get a list of TODOs that are not deleted.
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			page	query	int	false	"page number"
//	@Param			limit	query	int	false	"number of items per page"
//	@Success		200		{array}	todos.Todo
//	@Router			/todos [get]
func (c *TodoController) TodoListHandler(w http.ResponseWriter, req *http.Request) {
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

	todoList, err := c.service.TodoListService(page, limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("X-CSRF-Token", csrf.Token(req))
	json.NewEncoder(w).Encode(todoList)
	w.WriteHeader(http.StatusOK)
}

//	@Summary		Delete todo
//	@Description	delete todo by id
//	@Tags			todos
//	@Accept			json
//	@Param			id	path	int	true	"user id"
//	@Success		200
//	@Router			/todos/{id}/delete [post]
func (c *TodoController) DeleteTodoHandler(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if err := c.service.DeleteTodo(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
