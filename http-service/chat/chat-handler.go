package chat

import (
	"io"
	"log"

	"golang.org/x/net/websocket"
)

type chatHandlers struct{}

func NewChatHandlers() *chatHandlers {
	return &chatHandlers{}
}

func (h *chatHandlers) Connect(ws *websocket.Conn) {
	defer func() {
		err := ws.Close()
		if err != nil {
			log.Println("Error closing websocket:", err)
		}
	}()

	_, err := io.Copy(ws, ws)
	if err != nil {
		log.Println("Error copying data:", err)
	}
}
