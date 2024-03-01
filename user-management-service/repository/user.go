package repository

import (
	"errors"

	"gorm.io/gorm"
)

// this file sould be moved to and replaced with gRPC call to user-managment service
type UserRepository interface {
  // StoreUser operation
	Save(user User) (*User, error)
  // GetUser operation
	FindByUsername(name string) (*User, error)
  // GetUser operation
	FindUserByCreds(creds *UserCreds) (*User, error)
}

type UserPgRepository struct {
	db *gorm.DB
}

type UserCreds struct {
	Name     string
	Password string
}

func NewUserPgRepository() *UserPgRepository {
	return &UserPgRepository{
		db: UserDBConnect(),
	}
}

func (repo *UserPgRepository) Save(user User) (*User, error) {
	if repo.db == nil {
		return nil, errors.New("database not initialized")
	}

	if user.Name == "" {
		return nil, errors.New("User name cannot be empty")
	}

	savedUser, _ := repo.FindByUsername(user.Name)

	if savedUser != nil {
		return savedUser, nil
	}

	result := repo.db.Create(&user)

	if result.Error != nil {
		return nil, errors.New("error while saving user to db")
	}

	return &user, nil
}

func (repo *UserPgRepository) FindByUsername(name string) (*User, error) {
	if name == "" {
		return nil, errors.New("user name cannot be empty")
	}

	user := &User{}

	result := repo.db.Where("name = ?", name).First(user)

	if result.Error != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (repo *UserPgRepository) FindUserByCreds(creds *UserCreds) (*User, error) {
	user, err := repo.FindByUsername(creds.Name)

	if err != nil || user.Password != creds.Password {
		return nil, errors.New("wrong username or password")
	}

	return user, nil
}
