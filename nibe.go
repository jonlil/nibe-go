package nibe

import (
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
		ClientSecret:    os.Getenv("CLIENT_SECRET"),
		ClientID:        os.Getenv("CLIENT_ID"),
		OAuhRedirectURI: OAuthRedirectURI,
		Endpoint:        nibeUplinkAPI,
	}
}
