package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rafaelbmateus/go-bootstrap/api/handler"
	"github.com/stretchr/testify/assert"
)

func TestPlay(t *testing.T) {
	type test struct {
		url    string
		status int
	}

	tests := []test{
		{
			url:    "/playlist?city=florianópolis",
			status: 200,
		},
		{
			url:    "/playlist?lat=13.8333&lon=-88.9167",
			status: 200,
		},
		{
			url:    "/playlist",
			status: 400,
		},
		{
			url:    "/playlist?lat=13.8333",
			status: 400,
		},
		{
			url:    "/playlist?lon=-88.9167",
			status: 400,
		},
		{
			url:    "/playlist?city=floripa",
			status: 404,
		},
		{
			url:    "/playlist?city=9999",
			status: 404,
		},
	}

	for _, tc := range tests {
		r := gin.Default()
		r.GET("/playlist", handler.Play)

		req, _ := http.NewRequest(http.MethodGet, tc.url, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, tc.status, w.Code)
	}
}
