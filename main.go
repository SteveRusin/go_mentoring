package main

import (
	"log/slog"
	"net/http"

	"github.com/SteveRusin/go_mentoring/users"
)

func main() {
	http.HandleFunc("/user", users.HandleUser)
	http.HandleFunc("/user/login", users.HandleUserLogin)

	slog.Info("Server is listening on localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}
