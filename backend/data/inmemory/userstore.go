package inmemory

import (
	"errors"

	"github.com/DiscoFighter47/todo/backend/model"
)

// UserStore ..
type UserStore struct {
	repo map[string]*model.User
}

// NewUserStore ...
func NewUserStore() *UserStore {
	return &UserStore{
		repo: map[string]*model.User{},
	}
}

// AddUser ...
func (store *UserStore) AddUser(user *model.User) error {
	if _, ok := store.repo[user.ID]; ok {
		return errors.New("user already exists")
	}
	store.repo[user.ID] = user
	return nil
}

// GetUser ...
func (store *UserStore) GetUser(id string) (*model.User, error) {
	if user, found := store.repo[id]; found {
		return user, nil
	}
	return nil, errors.New("user not found")
}
