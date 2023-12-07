package service

import (
	"errors"
	"sync"
)

type UserStore interface {
	Save(user *User) error
	Find(username string) (*User, error)
	CreateUser(username, password, role string) error
}

type InMemoryUserStore struct {
	mutex sync.RWMutex
	users map[string]*User
}

func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		users: make(map[string]*User),
	}
}

func (store *InMemoryUserStore) Save(user *User) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.users[user.Username] != nil {
		return errors.New("ErrAlreadyExists")
	}

	store.users[user.Username] = user.Clone()
	return nil
}

func (store *InMemoryUserStore) Find(username string) (*User, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	user := store.users[username]
	if user == nil {
		return nil, nil
	}

	return user.Clone(), nil
}

func (store *InMemoryUserStore) CreateUser(username, password, role string) error {
	user, err := NewUser(username, password, role)
	if err != nil {
		return err
	}
	return store.Save(user)
}
