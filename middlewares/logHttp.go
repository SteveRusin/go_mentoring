package middlewares

import (
	"log/slog"
	"net/http"
)

func LogHttpMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		path := r.URL.Path

		slog.Info(method, "path", path)
		next.ServeHTTP(w, r)
	})
}
