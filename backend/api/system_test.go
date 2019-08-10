package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSystemCheck(t *testing.T) {
	api := NewAPI(nil, nil)
	t.Run("system check", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/system/check", nil)
		res := httptest.NewRecorder()
		api.handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.JSONEq(t, `{"data":{"message":"Hello Universe!"}}`, res.Body.String())
	})
}

func TestSystemPanic(t *testing.T) {
	api := NewAPI(nil, nil)
	t.Run("system panic", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/system/panic", nil)
		res := httptest.NewRecorder()
		api.handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.JSONEq(t, `{"error":{"title":"System Panic!","detail":"system test panic"}}`, res.Body.String())
	})
}

func TestSystemError(t *testing.T) {
	api := NewAPI(nil, nil)
	t.Run("system error", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/system/err", nil)
		res := httptest.NewRecorder()
		api.handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.JSONEq(t, `{"error":{"title":"Internal Server Error","detail":"runtime error: integer divide by zero"}}`, res.Body.String())
	})
}
