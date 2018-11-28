package nibe

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// PaginatedAPIResponse -
type PaginatedAPIResponse struct {
	Page         int
	ItemsPerPage int
	NumItems     int
	Objects      json.RawMessage
}

// API -
type API struct {
	AuthorizationToken string
	httpClient         *http.Client
	Endpoint           string
}

// NewAPI -
func NewAPI(authorizationToken string) *API {
	api := &API{
		AuthorizationToken: authorizationToken,
		httpClient:         &http.Client{},
		Endpoint:           nibeUplinkAPI,
	}
	return api
}

func (api *API) buildAuthorizationHeader() []string {
	return []string{
		"Authorization",
		"Bearer " + api.AuthorizationToken,
	}
}

func (api *API) getURL(endpoint string) string {
	return api.Endpoint + "/" + endpoint
}

// Get -
func (api *API) Get(url string, data interface{}) error {
	return api.request("GET", url, &data)
}

// request - Helper method for making requests
func (api *API) request(method string, url string, data *interface{}) error {
	req, err := http.NewRequest(strings.ToUpper(method), api.getURL(url), nil)

	authHeader := api.buildAuthorizationHeader()
	req.Header.Add(authHeader[0], authHeader[1])

	resp, err := api.httpClient.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}
	switch status := resp.StatusCode; status {
	case 401:
		return errors.New("Unauthenticated")
	case 200:
	default:
		return errors.New("Unhandled response StatusCode: " + string(status))
	}
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return readErr
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	return nil
}
