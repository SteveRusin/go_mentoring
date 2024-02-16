package users

import (
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type UserCreds struct {
	Name     string
	Password string
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: UserDBConnect(),
	}
}

func (repo *UserRepository) Save(user User) (*User, error) {
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

func (repo *UserRepository) FindByUsername(name string) (*User, error) {
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

func (repo *UserRepository) FindUserByCreds(creds *UserCreds) (*User, error) {
	user, err := repo.FindByUsername(creds.Name)

	if err != nil || user.Password != creds.Password {
		return nil, errors.New("wrong username or password")
	}

	return user, nil
}
