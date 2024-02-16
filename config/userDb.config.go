package config

import "os"

type userDbConfig struct {
  User string
  Password string
}

var userConfig *userDbConfig

func GetUserDbConfig() *userDbConfig {
  if userConfig == nil {
    userConfig = &userDbConfig{
      User: os.Getenv("POSTGRES_USER"),
      Password: os.Getenv("POSTGRES_PASSWORD"),
    }

    if userConfig.User == "" || userConfig.Password == "" {
      panic("POSTGRES_USER or POSTGRES_PASSWORD not defined")
    }
  }

  return userConfig
}
