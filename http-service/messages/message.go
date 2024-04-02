package messages

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Message struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`

	Text   string `bson:"text"`
	UserId string `bson:"user_id"`
}

type MessageRepository struct {
	db *mongo.Collection
}

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{
		db: GetMessagesCollection(),
	}
}

func (repo *MessageRepository) FindAllExcept(userId string) ([]Message, error) {
	filter := bson.D{{"user_id", bson.D{{"$ne", userId}}}}

	cursor, err := repo.db.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var results []Message
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("user %s does not have messages", userId)
	}
	return results, nil
}

func (repo *MessageRepository) FindAllByUser(userId string) ([]Message, error) {
	filter := bson.D{{"user_id", userId}}

	cursor, err := repo.db.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var results []Message
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("user %s does not have messages", userId)
	}
	return results, nil
}

func (repo *MessageRepository) Save(userId string, text string) error {
	if repo.db == nil {
		return errors.New("message db is not initialized")
	}

	message := Message{
		Text:   text,
		UserId: userId,
	}

	b, err := bson.Marshal(&message)
	if err != nil {
		return err
	}

	_, err = repo.db.InsertOne(context.TODO(), b)

	return err
}
