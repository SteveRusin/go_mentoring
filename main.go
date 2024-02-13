package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/SteveRusin/go_mentoring/users"

	"github.com/SteveRusin/go_mentoring/middlewares"
)

func main() {
	handleUser := http.HandlerFunc(users.HandleUser)
	handleUserLogin := http.HandlerFunc(users.HandleUserLogin)
	mux := http.NewServeMux()

	mux.Handle("/user", middlewares.CatchPanicMiddleware(middlewares.LogHttpMiddleware(handleUser)))

	mux.Handle("/user/login", middlewares.CatchPanicMiddleware(middlewares.LogHttpMiddleware(handleUserLogin)))

	slog.Info("Server is listening on localhost:8080")
	err := http.ListenAndServe("localhost:8080", mux)

	slog.Error(fmt.Sprint(err))
}

