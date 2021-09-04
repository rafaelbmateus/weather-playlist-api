package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rafaelbmateus/go-bootstrap/api/handler"
	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	route := gin.Default()
	route.GET("/", handler.Home)

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	route.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
