package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	t.Parallel()

	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())

	code, _, body := ts.get(t, "/health")
	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, `{"status": "ok"}`)
}
