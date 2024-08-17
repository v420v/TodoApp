package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/uptrace/bun"
	"github.com/v420v/go-api/internal/app/controller"
	"github.com/v420v/go-api/internal/app/service"
)

func NewRouter(db *bun.DB) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello world!\n"))
	}).Methods(http.MethodGet)

	service := service.NewService(db)
	todoController := controller.NewTodoController(service)

	r.HandleFunc("/todos", todoController.PostTodoHandler).Methods(http.MethodPost)
	r.HandleFunc("/todos", todoController.TodoListHandler).Methods(http.MethodGet) // TODO: add page number and length
	r.HandleFunc("/todos/{id:[0-9]+}/delete", todoController.DeleteTodoHandler).Methods(http.MethodPost)

	return r
}
