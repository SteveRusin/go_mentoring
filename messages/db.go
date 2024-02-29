package messages

import (
	"context"
	"fmt"

	"github.com/SteveRusin/go_mentoring/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

func MessagesDBConnect() *mongo.Client {
	if db != nil {
		return db
	}

	config := config.GetMessagesDbConfig()
	connectString := fmt.Sprintf("mongodb://%s:%s@localhost:27017/", config.User, config.Password)

	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectString))

	if err != nil {
		panic(err)
	}

	return db
}

func GetMessagesCollection() *mongo.Collection {
  db := MessagesDBConnect()

  return db.Database("go_mentoring").Collection("messages")
}
