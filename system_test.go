package nibe

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"log"
)

func TestGetSystems(t *testing.T) {
	api := NewAPI("mysecretaccesstoken")

	systems, err := api.GetSystems()
	assert.Equal(t, err, nil)
	for _, element := range *systems {
		log.Println(element)
	}
}
