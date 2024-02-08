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

type UserCreds struct {
  Name string
  Password string
}

func NewUserRepository() *UserRepository {
	var usersDb UserDb = make(map[string]User)

	return &UserRepository{
		db: usersDb,
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

	repo.db[user.Id] = user

	return &user, nil
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

func (repo *UserRepository) FindUserByCreds(creds *UserCreds) (*User, error) {
  user, err := repo.FindByUsername(creds.Name)

  if err != nil || user.Password != creds.Password {
    return nil, errors.New("Wrong username or password")
  }

  return user, nil
}
