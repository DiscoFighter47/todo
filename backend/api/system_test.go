package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSystemCheck(t *testing.T) {
	api := NewAPI(nil)
	t.Run("system check", func(t *testing.T) {
		res := httptest.NewRecorder()
		api.systemCheck(res, nil)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.JSONEq(t, `{"data":{"message":"Hello Universe!"}}`, res.Body.String())
	})
}

func TestSystemPanic(t *testing.T) {
	api := NewAPI(nil)
	t.Run("system panic", func(t *testing.T) {
		assert.Panics(t, func() { api.systemPanic(nil, nil) })
	})
}

func TestSystemError(t *testing.T) {
	api := NewAPI(nil)
	t.Run("system error", func(t *testing.T) {
		assert.Panics(t, func() { api.systemError(nil, nil) })
	})
}
