package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/SteveRusin/go_mentoring/messages"
	"github.com/SteveRusin/go_mentoring/users"
	_ "github.com/joho/godotenv/autoload" // read .env file

	"github.com/SteveRusin/go_mentoring/middlewares"
)

func main() {
  users.MigrateUsersDb()
  // messages.NewMessagRepository().Save("123", messages.Message{
  //   Id: "msgId",
  //   Text: "example",
  // })
  messages.NewMessagRepository().FindAll("1232")
	mux := http.NewServeMux()

	mux.Handle(
		"/user",
		middlewares.CatchPanicMiddleware(
			middlewares.LogHttpMiddleware(
				middlewares.ErrorMiddleware(
					users.HandleUser,
				),
			),
		),
	)

	mux.Handle(
		"/user/login",
		middlewares.CatchPanicMiddleware(
			middlewares.LogHttpMiddleware(
				middlewares.ErrorMiddleware(
					users.HandleUserLogin,
				),
			),
		),
	)

	slog.Info("Server is listening on localhost:8080")
	err := http.ListenAndServe("localhost:8080", mux)

	slog.Error(fmt.Sprint(err))
}
