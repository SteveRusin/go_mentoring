package messages

import (
	"errors"
	"fmt"
)

type MessageDb map[string][]Message

type Message struct {
	Id   string
	Text string
}

type MessageRepository struct {
	db MessageDb
}

func NewMessagRepository() *MessageRepository {
	var messageDb MessageDb = make(map[string][]Message)

	return &MessageRepository{
		db: messageDb,
	}
}

func (repo *MessageRepository) findAll(userId string) ([]Message, error) {
	messages, ok := repo.db[userId]

	if !ok {
		return nil, fmt.Errorf("user %s does not have messages", userId)
	}

	return messages, nil
}

func (repo *MessageRepository) Save(userId string, message Message) (error) {
  if repo.db == nil {
    return errors.New("message db is not initialized")
  }

  repo.db[userId] = append(repo.db[userId], message)

  return nil
}
