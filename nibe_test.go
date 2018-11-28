package nibe

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the API client being tested
	client *API

	// server is a test HTTP server used to provide mock API responses
	server *httptest.Server
)

func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewAPI("mysecrettoken")
	client.Endpoint = server.URL
}

func teardown() {
	server.Close()
}

func TestAPIClient_have_Endpoint(t *testing.T) {
	setup()
	assert.NotEqual(t, client.Endpoint, nil)
	teardown()
}

func TestRefreshToken_should_refresh(t *testing.T) {
	_, err := RefreshToken("myrefreshtoken")
	assert.Equal(t, nil, err, "Err should be nil")
}
