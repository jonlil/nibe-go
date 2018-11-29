package nibe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const parameterResponse string = `
[
  {"parameterId":40033,"name":"indoor_temperature","title":"room temperature","designation":"BT50","unit":"°C","displayValue":"22.6°C","rawValue":226},
  {"parameterId":40004,"name":"outdoor_temperature","title":"outdoor temp.","designation":"BT1","unit":"°C","displayValue":"-3.7°C","rawValue":-37}
]`

func TestGetParameters_outdoortemp(t *testing.T) {
	setup()

	mux.HandleFunc("/api/v1/systems/12345/parameters", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected method 'GET', got %s", r.Method)
		assert.Contains(t, r.URL.String(), "parameterIds=40004&parameterIds=40033", "Parameters should expand to query string")

		fmt.Fprintf(w, parameterResponse)
	})

	response, err := client.GetParameters(&System{
		SystemID: 12345,
	}, []string{
		"40004",
		"40033",
	})

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, response)
	assert.Equal(t, "indoor_temperature", response[0].Name)
	assert.Equal(t, "room temperature", response[0].Title)
	assert.Equal(t, "BT50", response[0].Designation)
	assert.Equal(t, "°C", response[0].Unit)

	teardown()
}
