package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/uptrace/bun"
	"github.com/v420v/go-api/internal/api/controller"
	"github.com/v420v/go-api/internal/api/middleware"
	"github.com/v420v/go-api/internal/infrastructure/repository"
	"github.com/v420v/go-api/internal/todo/app/service"
)

func NewRouter(db *bun.DB) *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.CORSMiddleware)

	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello world!\n"))
	}).Methods(http.MethodGet)

	repository := repository.NewTodoRepository(db)
	todoService := service.NewTodoApplicationService(repository)
	todoController := controller.NewTodoController(todoService)

	r.HandleFunc("/todos", todoController.GetList).Methods(http.MethodGet)
	r.HandleFunc("/todos", todoController.PostTodo).Methods(http.MethodPost)
	r.HandleFunc("/todos/{id}/delete", todoController.DeleteTodo).Methods(http.MethodPost)

	return r
}
