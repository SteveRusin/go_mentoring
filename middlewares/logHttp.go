package middlewares

import (
	"log/slog"
	"net/http"
	"time"
)

func LogHttpMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		path := r.URL.Path

		slog.Info(method, "path", path)
		start := time.Now()

		next.ServeHTTP(w, r)

		elapsed := time.Since(start)
		slog.Info(method, "path", path, "elapsed", elapsed)
	})
}
