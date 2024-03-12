package config

import "os"

type userDbConfig struct {
	Host     string
	User     string
	Password string
}

var userConfig *userDbConfig

func GetUserDbConfig() *userDbConfig {
	if userConfig == nil {
		userConfig = &userDbConfig{
			Host:     os.Getenv("POSTGRES_HOST"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
		}

		if userConfig.Host == "" || userConfig.User == "" || userConfig.Password == "" {
			panic("POSTGRES_HOST, POSTGRES_USER or POSTGRES_PASSWORD not defined")
		}
	}

	return userConfig
}
