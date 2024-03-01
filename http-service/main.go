package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/SteveRusin/go_mentoring/http-service/config"
	"github.com/SteveRusin/go_mentoring/http-service/users"
	_ "github.com/joho/godotenv/autoload" // read .env file

	"github.com/SteveRusin/go_mentoring/http-service/middlewares"
)

func main() {
	mux := http.NewServeMux()

	usersHandler := users.NewUserHandlers()

	mux.Handle(
		"/user",
		middlewares.CatchPanicMiddleware(
			middlewares.LogHttpMiddleware(
				middlewares.ErrorMiddleware(
					usersHandler.User,
				),
			),
		),
	)

	mux.Handle(
		"/user/login",
		middlewares.CatchPanicMiddleware(
			middlewares.LogHttpMiddleware(
				middlewares.ErrorMiddleware(
					usersHandler.UserLogin,
				),
			),
		),
	)

  host := fmt.Sprintf("%s:8080", config.GetAppConfig().Host)
	slog.Info(fmt.Sprintf("Server is listening on %s", host))
	err := http.ListenAndServe(host, mux)

	slog.Error(fmt.Sprint(err))
}
