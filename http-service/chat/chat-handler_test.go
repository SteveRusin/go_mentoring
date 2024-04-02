package chat

import (
	"context"
	"log"
	"net/http"
	"testing"

	"golang.org/x/net/websocket"
)

func TestShouldConnectToChat(t *testing.T) {
	handler := chatHandlers{
		connectionsByUser:  make(map[string]*websocket.Conn),
		messagesRepository: nil,
	}

	// how to mock this?
	conn, err := websocket.Dial("ws://localhost", "", "http://localhost")
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	defer conn.Close()

	req, err := http.NewRequest("GET", "/", nil)
	req = req.WithContext(context.WithValue(req.Context(), "user", "test"))

	if err != nil {
		log.Println(err)
		t.Fail()
	}

	handler.Connect(conn)

	log.Println("Test passed")
}
