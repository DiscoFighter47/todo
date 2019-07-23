package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kinbiko/jsonassert"
)

func TestSystemCheck(t *testing.T) {
	api := NewAPI()
	req := httptest.NewRequest("GET", "/system/check", nil)
	res := httptest.NewRecorder()
	api.systemCheck(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	jsonassert.New(t).Assertf(res.Body.String(), `{"Code":200,"data":{"message":"Hello Universe!"}}`)
}
