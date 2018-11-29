package nibe

import (
	"fmt"
	"net/url"
)

// Parameter -
type Parameter struct {
	ParameterID  uint32
	Name         string
	Title        string
	Designation  string
	Unit         string
	DisplayValue string
	RawValue     int
}

// GetParameters -
func (api *API) GetParameters(system *System, parameters []string) ([]Parameter, error) {
	var response []Parameter
	params := url.Values{}
	for _, v := range parameters {
		params.Add("parameterIds", v)
	}

	err := api.Get(fmt.Sprintf("/api/v1/systems/%d/parameters?%s",
		system.SystemID,
		params.Encode(),
	), &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
