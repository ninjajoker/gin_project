package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestHelloRoute(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET","/hello",nil)
	r.ServeHTTP(w,req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "world" , w.Body.String())
}
