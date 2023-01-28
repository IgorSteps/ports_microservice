package client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var expectedPort = `{"name":"Ajman","city":"Ajman","country":"United Arab Emirates","alias":[],"regions":[],"coordinates":[55.5136433,25.4052165],"province":"Ajman","timezone":"Asia/Dubai","unlocs":["AEAJM"],"code":"52000"}`

// func TestGetMany(t *testing.T) {
// 	router := gin.Default()
// 	router.GET("/ports", GetMultiple)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/ports", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// 	assert.Equal(t, "pong", w.Body.String())
// }

func TestGetSingle(t *testing.T) {
	router := gin.Default()
	router.GET("/port/:port_code", GetSingle)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/port/AEAJM", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expectedPort, w.Body.String())
}
