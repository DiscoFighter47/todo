package inmemory

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DiscoFighter47/todo/backend/data"
	"github.com/DiscoFighter47/todo/backend/model"
)

func TestUserStore(t *testing.T) {
	store := NewUserStore()

	user := &model.User{
		ID:       "DiscoFighter47",
		Name:     "Zahid Al Tair",
		Password: "password",
	}

	_, err := store.GetUser(user.ID)
	assert.Equal(t, data.ErrUserNotFound, err)

	err = store.AddUser(user)
	assert.Nil(t, err)

	u, err := store.GetUser(user.ID)
	assert.Nil(t, err)
	assert.Equal(t, user, u)

	err = store.AddUser(user)
	assert.Equal(t, data.ErrUserAlreadyExists, err)
}
