package nibe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSystems(t *testing.T) {
	api := NewAPI("mysecretaccesstoken")

	_, err := api.GetSystems()
	assert.Equal(t, nil, err)
}
