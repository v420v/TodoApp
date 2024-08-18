package middleware

import (
	"net/http"

	"github.com/charmbracelet/log"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "http://127.0.0.1:5173" {
			origin = "http://127.0.0.1:80"
		}
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, DELETE, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("AllowCredentials", "true")

		if r.Method == "OPTIONS" {
			log.Info("pre-flight request")
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
