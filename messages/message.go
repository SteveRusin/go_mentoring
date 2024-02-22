package messages

import (
	"errors"
	"fmt"
	"sync"
)

type MessageDb map[string][]Message

var mutex *sync.Mutex = &sync.Mutex{}

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
	mutex.Lock()
	defer mutex.Unlock()

	messages, ok := repo.db[userId]

	if !ok {
		return nil, fmt.Errorf("user %s does not have messages", userId)
	}

	return messages, nil
}

func (repo *MessageRepository) Save(userId string, message Message) error {
	if repo.db == nil {
		return errors.New("message db is not initialized")
	}

	mutex.Lock()
	defer mutex.Unlock()

	repo.db[userId] = append(repo.db[userId], message)

	return nil
}
