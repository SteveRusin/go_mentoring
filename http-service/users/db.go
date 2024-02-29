package users

import (
	"fmt"
	"log/slog"

	"github.com/SteveRusin/go_mentoring/http-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Id       string `gorm:"primaryKey"`
	Name     string `gorm:"uniqueIndex"` // make unique just for the sake of error handling example
	Password string
}

var db *gorm.DB

func UserDBConnect() *gorm.DB {
	if db != nil {
		return db
	}
	config := config.GetUserDbConfig()

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s  user=%s password=%s", config.Host, config.User, config.Password),
	}))
	if err != nil {
		panic("Unable to connect to users database")
	}

	return db
}

func MigrateUsersDb() {
	slog.Info("Starting user migrations")
	db := UserDBConnect()

	db.AutoMigrate(&User{})

	slog.Info("User migrations are completed")
}
