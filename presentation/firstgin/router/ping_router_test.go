package router

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingAndSuccess(t *testing.T) {
	// object gin.Engine
	r := SetupEngine(Ping)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestPingWithNameAndSuccess(t *testing.T) {
	r := SetupEngine(Ping, PingAction)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping/bilal", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong bilal", w.Body.String())
}

func TestPingWithNameAnsAsterisksParamAndSuccess(t *testing.T) {
	r := SetupEngine(Ping, PingAction)

	w := httptest.NewRecorder()
	// /ping/:name/*action
	req, _ := http.NewRequest("GET", "/ping/bilal/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong bilal the action /bike", w.Body.String())
}
