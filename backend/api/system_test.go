package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"
)

func TestSystemCheck(t *testing.T) {
	api := NewAPI(nil)
	t.Run("system check", func(t *testing.T) {
		res := httptest.NewRecorder()
		api.systemCheck(res, nil)
		assert.Equal(t, http.StatusOK, res.Code)
		jsonassert.New(t).Assertf(res.Body.String(), `{"data":{"message":"Hello Universe!"}}`)
	})
}

func TestSystemPanic(t *testing.T) {
	api := NewAPI(nil)
	t.Run("system panic", func(t *testing.T) {
		assert.Panics(t, func() { api.systemPanic(nil, nil) })
	})
}
