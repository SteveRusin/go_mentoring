package middlewares

import (
	"log/slog"
	"net/http"
)

type HttpError struct {
	statusCode uint16
	body       string
}

type handerFunc func(w http.ResponseWriter, r *http.Request) *HttpError

func ErrorMiddleware(next handerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := next(w, r); err != nil {
			slog.Error(err.body, "statusCode", err.statusCode)

			w.WriteHeader(int(err.statusCode))
			w.Write([]byte(err.body))
			w.Write([]byte("\n"))
		}
	})
}

func NewNotImplementedError() *HttpError {
	return &HttpError{
		statusCode: http.StatusNotImplemented,
		body:       "Method not implemented",
	}
}

func NewBadRequestError() *HttpError {
	return &HttpError{
		statusCode: http.StatusBadRequest,
		body:       "Bad request",
	}
}
