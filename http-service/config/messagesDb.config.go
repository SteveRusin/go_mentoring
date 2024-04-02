package config

import "os"

type messagesDbConfig struct {
	User     string
	Password string
}

var messagesConfig *messagesDbConfig

func GetMessagesDbConfig() *messagesDbConfig {
	if messagesConfig == nil {
		messagesConfig = &messagesDbConfig{
			User:     os.Getenv("MONGO_USER"),
			Password: os.Getenv("MONGO_PASSWORD"),
		}

		if messagesConfig.User == "" || messagesConfig.Password == "" {
			panic("MONGO_USER or MONGO_PASSWORD not defined")
		}
	}

	return messagesConfig
}
