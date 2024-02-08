package main

import (
	"fmt"
	"net/http"

	"github.com/SteveRusin/go_mentoring/users"
)

func main() {
	http.HandleFunc("/user", users.HandleUser)
	http.HandleFunc("/user/login", users.HandleUserLogin)

	fmt.Println("Server is listening on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
