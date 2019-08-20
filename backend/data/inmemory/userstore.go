package inmemory

import (
	"sync"

	"github.com/DiscoFighter47/todo/backend/data"
	"github.com/DiscoFighter47/todo/backend/model"
)

// UserStore ..
type UserStore struct {
	repo sync.Map
}

// NewUserStore ...
func NewUserStore() *UserStore {
	return &UserStore{
		repo: sync.Map{},
	}
}

// AddUser ...
func (store *UserStore) AddUser(user *model.User) error {
	if _, ok := store.repo.Load(user.ID); ok {
		return data.ErrUserAlreadyExists
	}
	store.repo.Store(user.ID, user)
	return nil
}

// GetUser ...
func (store *UserStore) GetUser(id string) (*model.User, error) {
	if user, found := store.repo.Load(id); found {
		return user.(*model.User), nil
	}
	return nil, data.ErrUserNotFound
}
