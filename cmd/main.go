package main

import (
	"context"
	"errors"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmbracelet/log"
	"github.com/gorilla/mux"
	"github.com/v420v/go-api/internal/app/router"
	"github.com/v420v/go-api/internal/db"
)

const (
	PORT = ":8080"
)

func newServer(r *mux.Router) *http.Server {
	return &http.Server{
		Addr:              PORT,
		Handler:           r,
		ReadHeaderTimeout: 20 * time.Second,
	}
}

//	@title						Todo REST API
//	@version					1.0
//	@description				A simple todo rest api
//	@contact.email				ibuki420v@gmail.com
//	@license.name				MIT LICENSE
//	@license.url				https://github.com/v420v/TodoApp/blob/main/LICENSE.md
//	@host						127.0.0.1:8080
//	@BasePath					/
//	@externalDocs.description	GitHub
//	@externalDocs.url			https://github.com/v420v/TodoApp
func main() {
	db, err := db.ConnectDB()
	if err != nil {
		log.Errorf("failed to connect to database: %v", err)
		return
	}

	r := router.NewRouter(db)
	srv := newServer(r)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Infof("server starting at port %s", srv.Addr)
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("server error: %v", err)
		}
		log.Info("server stopped serving new connections")
	}()

	<-ctx.Done()
	log.Info("cached signal interrupt")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Errorf("server shutdown error: %v", err)
	}

	log.Info("server graceful shutdown complete")
}
