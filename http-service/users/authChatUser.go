package users

import (
	"log"
	"net/http"

	"github.com/SteveRusin/go_mentoring/http-service/middlewares"
	"golang.org/x/net/websocket"
)

type chatHandler = func(ws *websocket.Conn)

func AuthChatUser(handler chatHandler) middlewares.HanderFunc {
	db := newUsersActiveDb()
	return func(w http.ResponseWriter, r *http.Request) *middlewares.HttpError {
		token := r.URL.Query().Get("token")

		if token == "" {
			log.Println("No token provided")
			return middlewares.NewUnauthorizedError()
		}

		userName := db.GetUserByToken(token)

		if userName == "" {
			log.Println("No user found for token")
			return middlewares.NewUnauthorizedError()
		}

		db.MarkUserAsActive(userName)

		defer func() {
			db.RemoveUser(userName)
		}()

		db.RevokeToken(token)

		// how to pass active user data to the websocket handler?
		websocket.Handler(handler).ServeHTTP(w, r)

		return nil
	}
}
