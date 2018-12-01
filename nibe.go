package nibe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
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
	Token        string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

func (at *AccessToken) Refresh() error {
	nat, err := RefreshToken(at.RefreshToken)
	if err != nil {
		return err
	}

	at.Token = nat.Token
	at.RefreshToken = nat.RefreshToken
	at.Scope = nat.Scope
	at.TokenType = nat.TokenType
	at.ExpiresIn = nat.ExpiresIn

	return nil
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

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Failed refreshing token: %d", resp.StatusCode)
	}

	tokenData := &AccessToken{}
	err := json.Unmarshal(body, &tokenData)
	if err != nil {
		return nil, err
	}

	return tokenData, nil
}
