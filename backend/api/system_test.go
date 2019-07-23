package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSystemCheck(t *testing.T) {
	api := NewAPI()
	req := httptest.NewRequest("GET", "/system/check", nil)
	res := httptest.NewRecorder()
	api.systemCheck(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "Hello Universe!", res.Body.String())
}
