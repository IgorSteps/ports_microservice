package client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetMany_Successfull(t *testing.T) {
	router := gin.Default()
	router.GET("/ports", GetMultiple)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ports", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
