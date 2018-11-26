package nibe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthorizationHeader(t *testing.T) {
	api := NewAPI("averysecrettoken")
	assert.Equal(t, []string{"Authorization", "Bearer averysecrettoken"}, api.buildAuthorizationHeader())
}
