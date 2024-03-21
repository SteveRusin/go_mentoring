package chat

import (
	"log"

	"golang.org/x/net/websocket"
)

type chatHandlers struct {
	connectionsByUser map[string]*websocket.Conn
}

func NewChatHandlers() *chatHandlers {
	return &chatHandlers{
		connectionsByUser: make(map[string]*websocket.Conn),
	}
}

func (h *chatHandlers) Connect(ws *websocket.Conn) {
	// ask how this casting works
	user := ws.Request().Context().Value("user").(string)
	h.connectionsByUser[user] = ws

	defer func() {
		err := ws.Close()
		delete(h.connectionsByUser, user)
		if err != nil {
			log.Println("Error closing websocket:", err)
		}
	}()

	for {
		var message string
		err := websocket.Message.Receive(ws, &message)
		if err != nil {
			log.Println("Error receiving message:", err)
			break
		}
		h.fanOutMessage(user, message)
	}
}

func (h *chatHandlers) fanOutMessage(user string, message string) {
	for u, ws := range h.connectionsByUser {
		if u != user {
      go h.sendMessageTo(ws, user, message)
		}
	}
}

func (h *chatHandlers) sendMessageTo(ws *websocket.Conn, user string, message string) {
  err := websocket.Message.Send(ws, message)
  if err != nil {
    log.Println("Error sending message:", err)
  }
}
