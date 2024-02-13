package middlewares

import (
	"fmt"
	"log/slog"
	"net/http"
)

func CatchPanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				slog.Error(fmt.Sprint(err))

				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal server error\n"))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
