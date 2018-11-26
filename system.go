package nibe

import (
	"encoding/json"
	"time"
)

type Address struct {
	StreetAddress  string `json:"addressLine1"`
	StreetAddress2 string `json:"addressLine2"`
	PostalCode     string
	City           string
	Region         string
	Country        string
}

// System -
type System struct {
	SystemID         int
	Name             string
	ProductName      string
	SecurityLevel    string
	SerialNumber     string
	ConnectionStatus string
	HasAlarmed       bool
	LastActivityDate time.Time
	Address          Address
}

// GetSystems -
func (api *API) GetSystems() (*[]System, error) {
	response := &PaginatedAPIResponse{}
	err := api.Get("/api/v1/systems", response)
	if err != nil {
		return nil, err
	}

	systems := []System{}
	if err := json.Unmarshal(response.Objects, &systems); err != nil {
		return nil, err
	}

	return &systems, nil
}
