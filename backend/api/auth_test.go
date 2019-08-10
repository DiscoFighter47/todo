package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	gauth "github.com/DiscoFighter47/gAuth"

	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"

	"github.com/DiscoFighter47/todo/backend/data/inmemory"
	"github.com/DiscoFighter47/todo/backend/model"
)

func TestAuthSignUp(t *testing.T) {
	api := NewAPI(inmemory.NewDatastore(), nil)

	testData := []struct {
		des  string
		body string
		code int
		res  string
	}{
		{
			des:  "sign up",
			body: `{"id": "DiscoFighter47","name": "Zahid Al Tair","password": "password"}`,
			code: http.StatusOK,
			res:  `{"data":{"id":"DiscoFighter47","name":"Zahid Al Tair"}}`,
		},
		{
			des:  "sign up duplicate",
			body: `{"id": "DiscoFighter47","name": "Zahid Al Tair","password": "password"}`,
			code: http.StatusInternalServerError,
			res:  `{"error":{"title":"Unable To Add User","detail":"User already exists","tags":["DiscoFighter47"]}}`,
		},
		{
			des:  "sign up invalid",
			body: `{}`,
			code: http.StatusBadRequest,
			res:  `{"error":{"title":"Invalid Request Body","detail":{"id":"required","name":"required","password":"required"}}}`,
		},
		{
			des:  "sign up faulty body",
			body: ``,
			code: http.StatusUnprocessableEntity,
			res:  `{"error":{"title":"Unable To Parse Body","detail":"EOF"}}`,
		},
	}

	for _, td := range testData {
		t.Run(td.des, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/auth/signup", bytes.NewReader([]byte(td.body)))
			res := httptest.NewRecorder()
			api.handler.ServeHTTP(res, req)
			assert.Equal(t, td.code, res.Code)
			assert.JSONEq(t, td.res, res.Body.String())
		})
	}
}

func TestAuthSignIn(t *testing.T) {
	store := inmemory.NewDatastore()
	store.AddUser(&model.User{
		ID:       "DiscoFighter47",
		Name:     "Zahid Al tair",
		Password: "password",
	})
	api := NewAPI(store, gauth.NewAuth("secret", 5*time.Second))

	testData := []struct {
		des  string
		body string
		code int
		res  string
	}{
		{
			des:  "sign in",
			body: `{"id": "DiscoFighter47","password": "password"}`,
			code: http.StatusOK,
			res:  `{"data":{"id":"DiscoFighter47","token":"<<PRESENCE>>"}}`,
		},
		{
			des:  "sign in wrong password",
			body: `{"id": "DiscoFighter","password": "password"}`,
			code: http.StatusBadRequest,
			res:  `{"error":{"title":"Unable To Find User","detail":"User not found","tags":["DiscoFighter"]}}`,
		},
		{
			des:  "sign in wrong password",
			body: `{"id": "DiscoFighter47","password": "pass"}`,
			code: http.StatusBadRequest,
			res:  `{"error":{"title":"Incorrect Password","detail":"password dosen't match","tags":["DiscoFighter47"]}}`,
		},
		{
			des:  "sign in invalid body",
			body: `{}`,
			code: http.StatusBadRequest,
			res:  `{"error":{"title":"Invalid Request Body","detail":{"id":"required","password":"required"}}}`,
		},
		{
			des:  "sign up faulty body",
			body: ``,
			code: http.StatusUnprocessableEntity,
			res:  `{"error":{"title":"Unable To Parse Body","detail":"EOF"}}`,
		},
	}

	for _, td := range testData {
		t.Run(td.des, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/auth/signin", bytes.NewReader([]byte(td.body)))
			res := httptest.NewRecorder()
			api.handler.ServeHTTP(res, req)
			assert.Equal(t, td.code, td.code)
			jsonassert.New(t).Assertf(res.Body.String(), td.res)
		})
	}
}

func TestAuthCheck(t *testing.T) {
	auth := gauth.NewAuth("secret", 1*time.Second)
	token, _ := auth.GenerateToken("user1")
	api := NewAPI(nil, auth)

	testData := []struct {
		des   string
		token string
		code  int
		res   string
	}{
		{
			des:   "auth check",
			token: "Bearer " + token,
			code:  http.StatusOK,
			res:   `{"data":{"message":"Hello Secret Universe! Welcome user1"}}`,
		},
		{
			des:  "no token",
			code: http.StatusUnauthorized,
			res:  `{"error":{"title":"Authorization Required","detail":"token not found"}}`,
		},
	}

	for _, td := range testData {
		t.Run(td.des, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/auth/check", nil)
			req.Header.Add("Authorization", td.token)
			res := httptest.NewRecorder()
			api.handler.ServeHTTP(res, req)
			assert.Equal(t, td.code, res.Code)
			assert.JSONEq(t, td.res, res.Body.String())
		})
	}
}
