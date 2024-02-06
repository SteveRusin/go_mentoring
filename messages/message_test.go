package messages

import (
	"fmt"
	"testing"
)

var messageRepository MessageRepository

func setUp() {
	messageRepository = *NewMessagRepository()
}

func saveMockMessageToDb() {
	mockMessage := Message{
		Id:   "1",
		Text: "Mock text",
	}

	messageRepository.Save("111", mockMessage)
}

func TestShouldSaveMessage(t *testing.T) {
	setUp()
	saveMockMessageToDb()
	savedMessages := messageRepository.db["111"]

	if savedMessages[0].Id != "1" {
		t.Fatal("Message not saved in db")
	}
}

func TestShouldFindAllMessages(t *testing.T) {
	setUp()
	saveMockMessageToDb()

	messages, _ := messageRepository.findAll("111")

	if messages[0].Id != "1" {
		t.Fatal("Message not saved in db")
	}
}

func TestShouldReturnErrorIfMessageNotFound(t *testing.T) {
	setUp()
	_, err := messageRepository.findAll("111")

	if fmt.Sprint(err) != "user 111 does not have messages" {
		t.Fatal("Should return an error when message is not found")
	}
}
