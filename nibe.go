package nibe

import (
	"os"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

const nibeUplinkAPI = "https://api.nibeuplink.com"

// Credentials - struct for accessing credentials
type Credentials struct {
	ClientSecret    string
	ClientID        string
	OAuhRedirectURI string
	Endpoint        string
}

// NewCredentials - Helper for settings common values
func NewCredentials(OAuthRedirectURI string) *Credentials {
	return &Credentials{
		ClientSecret:    os.Getenv("NIBE_APP_SECRET"),
		ClientID:        os.Getenv("NIBE_APP_ID"),
		OAuhRedirectURI: OAuthRedirectURI,
		Endpoint:        nibeUplinkAPI,
	}
}

// AccessToken -
type AccessToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

// RefreshTokenRequest -
type RefreshTokenRequest struct {
	GrantType    string
	RefreshToken string
}

// RefreshToken -
func RefreshToken(refreshToken string) (*AccessToken, error) {
	resp, _ := http.PostForm(
		fmt.Sprintf("%s/oauth/token", nibeUplinkAPI),
		url.Values{
			"grant_type":    {"refresh_token"},
			"client_id":     {os.Getenv("NIBE_APP_ID")},
			"client_secret": {os.Getenv("NIBE_APP_SECRET")},
			"refresh_token": {refreshToken},
		})
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}

	tokenData := &AccessToken{}
	err := json.Unmarshal(body, &tokenData); if err != nil {
		return nil, err
	}

	return tokenData, nil
}
