package config

import "os"

type userDbConfig struct {
  User string
  Password string
}

var config *userDbConfig

func GetUserDbConfig() *userDbConfig {
  if config == nil {
    config = &userDbConfig{
      User: os.Getenv("POSTGRES_USER"),
      Password: os.Getenv("POSTGRES_PASSWORD"),
    }

    if config.User == "" || config.Password == "" {
      panic("POSTGRES_USER or POSTGRES_PASSWORD not defined")
    }
  }

  return config
}
