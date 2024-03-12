package config

import "os"

type userServer struct {
	Host string
	Port string
}

var userServerConfig *userServer

func GetUserServerConfig() *userServer {
	if userServerConfig == nil {
		userServerConfig = &userServer{
			Host: os.Getenv("USER_SERVER_HOST"),
			Port: os.Getenv("USER_SERVER_PORT"),
		}

		if userServerConfig.Host == "" || userServerConfig.Port == "" {
			panic("USER_SERVER_HOST and USER_SERVER_PORT must be set")
		}
	}

	return userServerConfig
}
