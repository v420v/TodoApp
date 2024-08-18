package router

import (
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/uptrace/bun"
	"github.com/v420v/go-api/internal/app/controller"
	"github.com/v420v/go-api/internal/app/middleware"
	"github.com/v420v/go-api/internal/app/service"
)

func NewRouter(db *bun.DB) *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.CORSMiddleware)
	csrfMiddleware := csrf.Protect(
		[]byte("32-byte-long-auth-key"), // TODO: 32-byte-long-auth-key
		csrf.Secure(false),              // TODO: set true
	)
	r.Use(csrfMiddleware)

	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello world!\n"))
	}).Methods(http.MethodGet)

	service := service.NewService(db)
	todoController := controller.NewTodoController(service)

	r.HandleFunc("/todos", todoController.PostTodoHandler).Methods(http.MethodPost)
	r.HandleFunc("/todos", todoController.TodoListHandler).Methods(http.MethodGet)
	r.HandleFunc("/todos/{id:[0-9]+}/delete", todoController.DeleteTodoHandler).Methods(http.MethodPost)

	return r
}
