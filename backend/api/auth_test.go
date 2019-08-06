package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"

	"github.com/DiscoFighter47/todo/backend/data/inmemory"
	"github.com/DiscoFighter47/todo/backend/model"
)

func TestAuthSignUp(t *testing.T) {
	api := NewAPI(inmemory.NewDatastore())

	t.Run("sign up", func(t *testing.T) {
		body := `{"id": "DiscoFighter47","name": "Zahid Al Tair","password": "password"}`
		req := httptest.NewRequest("POST", "/auth/signup", bytes.NewReader([]byte(body)))
		res := httptest.NewRecorder()
		api.authSignUp(res, req)
		assert.Equal(t, res.Code, http.StatusOK)
		assert.JSONEq(t, `{"data":{"id":"DiscoFighter47","name":"Zahid Al Tair"}}`, res.Body.String())
	})

	t.Run("sign up duplicate", func(t *testing.T) {
		body := `{"id": "DiscoFighter47","name": "Zahid Al Tair","password": "password"}`
		req := httptest.NewRequest("POST", "/auth/signup", bytes.NewReader([]byte(body)))
		res := httptest.NewRecorder()
		assert.Panics(t, func() { api.authSignUp(res, req) })
	})

	t.Run("sign up faulty body", func(t *testing.T) {
		body := `{}`
		req := httptest.NewRequest("POST", "/auth/signup", bytes.NewReader([]byte(body)))
		res := httptest.NewRecorder()
		assert.Panics(t, func() { api.authSignUp(res, req) })
	})

	t.Run("sign up faulty body", func(t *testing.T) {
		body := ``
		req := httptest.NewRequest("POST", "/auth/signup", bytes.NewReader([]byte(body)))
		res := httptest.NewRecorder()
		assert.Panics(t, func() { api.authSignUp(res, req) })
	})
}

func TestAuthSignIn(t *testing.T) {
	store := inmemory.NewDatastore()
	store.AddUser(&model.User{
		ID:       "DiscoFighter47",
		Name:     "Zahid Al tair",
		Password: "password",
	})
	api := NewAPI(store)

	t.Run("sign in", func(t *testing.T) {
		body := `{"id": "DiscoFighter47","password": "password"}`
		req := httptest.NewRequest("POST", "/auth/signin", bytes.NewReader([]byte(body)))
		res := httptest.NewRecorder()
		api.authSignIn(res, req)
		assert.Equal(t, res.Code, http.StatusOK)
		jsonassert.New(t).Assertf(res.Body.String(), `{"data":{"id":"DiscoFighter47","token":"<<PRESENCE>>"}}`)
	})

	t.Run("sign in", func(t *testing.T) {
		body := `{"id": "DiscoFighter47","password": "pass"}`
		req := httptest.NewRequest("POST", "/auth/signin", bytes.NewReader([]byte(body)))
		res := httptest.NewRecorder()
		assert.Panics(t, func() { api.authSignIn(res, req) })
	})
}
