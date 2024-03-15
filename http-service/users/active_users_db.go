package users

import "sync"

type activeUsers struct {
	m      *sync.Mutex
	users  map[string]bool
	tokens map[string]string
}

var activeUsersDb *activeUsers

func newUsersActiveDb() *activeUsers {
	if activeUsersDb == nil {
		activeUsersDb = &activeUsers{
			m:      &sync.Mutex{},
			users:  make(map[string]bool),
			tokens: make(map[string]string),
		}
	}
	return activeUsersDb
}

func (db *activeUsers) AddToken(username, token string) {
	db.m.Lock()
	defer db.m.Unlock()

	db.tokens[token] = username
}

func (db *activeUsers) GetUserByToken(token string) string {
	db.m.Lock()
	defer db.m.Unlock()

	return db.tokens[token]
}

func (db *activeUsers) RevokeToken(token string) {
	db.m.Lock()
	defer db.m.Unlock()

	delete(db.tokens, token)
}

func (db *activeUsers) MarkUserAsActive(username string) {
	db.m.Lock()
	defer db.m.Unlock()

	db.users[username] = true
}

func (db *activeUsers) RemoveUser(username string) {
	db.m.Lock()
	defer db.m.Unlock()

	delete(db.users, username)
}

func (db *activeUsers) GetActiveUsers() []string {
	db.m.Lock()
	defer db.m.Unlock()

	users := make([]string, 0, len(db.users))
	for user := range db.users {
		users = append(users, user)
	}
	return users
}
