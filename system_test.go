package nibe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSystems(t *testing.T) {
	api := NewAPI("invalidtoken")

	_, err := api.GetSystems()
	assert.Equal(t, err, nil)
}
