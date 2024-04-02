package config

import "os"

type appConfig struct {
	Host string
}

var config *appConfig

func GetAppConfig() *appConfig {
	if config == nil {
		config = &appConfig{
			Host: os.Getenv("HOST"),
		}

		if config.Host == "" {
			config.Host = "0.0.0.0"
		}
	}

	return config
}
