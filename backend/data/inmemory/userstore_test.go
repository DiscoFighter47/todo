package inmemory

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DiscoFighter47/todo/backend/model"
)

func TestUserStore(t *testing.T) {
	store := NewUserStore()

	user := &model.User{
		ID:       "DiscoFighter47",
		Name:     "Zahid Al Tair",
		Password: "password",
	}

	err := store.AddUser(user)
	assert.Nil(t, err)

	u, err := store.GetUser(user.ID)
	assert.Nil(t, err)
	assert.Equal(t, user, u)

	err = store.AddUser(user)
	assert.NotNil(t, err)

	u, err = store.GetUser("hello universe!")
	assert.NotNil(t, err)
	assert.Nil(t, u)
}
