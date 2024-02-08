package users

import "errors"

type UserDb map[string]User

type User struct {
	Id       string
	Name     string
	Password string
}

type UserRepository struct {
	db UserDb
}

func NewUserRepository() *UserRepository {
	var usersDb UserDb = make(map[string]User)

	return &UserRepository{
		db: usersDb,
	}
}

func (repo *UserRepository) Save(user User) (*UserRepository, error) {
  if repo.db == nil {
    return nil, errors.New("database not initialized")
  }

	repo.db[user.Id] = user

	return repo, nil
}

func (repo *UserRepository) FindByUsername(name string) (*User, error) {
	if name == "" {
		return nil, errors.New("User name cannot be empty")
	}

	for _, v := range repo.db {
		if v.Name == name {
			return &v, nil
		}
	}

	return nil, errors.New("User not found")
}
