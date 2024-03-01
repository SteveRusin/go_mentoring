package config

import "os"

type appConfig struct {
	Host string
	Port string
}

var config *appConfig

func GetAppConfig() *appConfig {
	if config == nil {
		config = &appConfig{
			Host: os.Getenv("HOST"),
      Port: os.Getenv("PORT"),
		}

		if config.Host == "" {
			config.Host = "0.0.0.0"
		}

    if config.Port == "" {
      config.Port = "3001"
    }
	}

	return config
}
