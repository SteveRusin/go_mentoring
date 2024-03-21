package chat

import (
	"log"

	"github.com/SteveRusin/go_mentoring/http-service/messages"
	"golang.org/x/net/websocket"
)

type chatHandlers struct {
	connectionsByUser  map[string]*websocket.Conn
	messagesRepository *messages.MessageRepository
}

func NewChatHandlers() *chatHandlers {
	return &chatHandlers{
		connectionsByUser:  make(map[string]*websocket.Conn),
		messagesRepository: messages.NewMessageRepository(),
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

  h.sendUnreadMesages(user)

	for {
		err := h.receiveMessage(ws, user)
		if err != nil {
			break
		}
	}
}

func (h *chatHandlers) sendUnreadMesages(user string) {
  messages, err := h.messagesRepository.FindAllExcept(user)
  if err != nil {
    log.Println("Error finding messages:", err)
    return
  }

  ws := h.connectionsByUser[user]
  for _, message := range messages {
    go h.sendMessageTo(ws, message.UserId, message.Text)
    // mark as read
    // what would be a good design for this?
  }
}

func (h *chatHandlers) receiveMessage(ws *websocket.Conn, user string) error {
	var message string
	err := websocket.Message.Receive(ws, &message)
	if err != nil {
		log.Println("Error receiving message:", err)
		return err
	}

  // is this okay???
	go h.fanOutMessage(user, message)
  go func() {
    err := h.messagesRepository.Save(user, message)

    if err != nil {
      log.Println("Error saving message:", err)
    }
  }()

  return nil
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
