package chat

import (
	"io"

	"golang.org/x/net/websocket"
)

type chatHandlers struct{}

func NewChatHandlers() *chatHandlers {
	return &chatHandlers{}
}

func (h *chatHandlers) Connect(ws *websocket.Conn) {
	defer ws.Close()
	io.Copy(ws, ws)
}
