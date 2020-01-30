package client

import (
	"context"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// Client is an interface which all client types abide by. It guarantees operations around
// credentials; primarily getting, initializing, and refreshing.
type Client interface {
	Credentials() *RequestCredentials

	// InitializeCredentials should make the initial requests necessary to establish the first set of
	// authentication credentials within the Client.
	InitializeCredentials(context.Context) error

	// RefreshCredentials should initiate an internal refresh of the request credentials inside this
	// client. This refresh should, whenever possible, check the
	// RequestCredentials.AccessTokenExpiresAt field to determine whether it should actually refresh
	// the credentials or if the credentials are still valid.
	RefreshCredentials(context.Context) error

	// HTTPClient returns the `*http.Client` to use for this Client.
	HTTPClient() *http.Client
}

// RequestCredentials stores all the information necessary to authenticate a request with the
// Microsoft GraphAPI
type RequestCredentials struct {
	AccessToken          string
	AccessTokenExpiresAt time.Time
	AccessTokenUpdating  sync.Mutex
}

// ConsentURL builds the consent URL needed to add this application to a target Azure domain.
func ConsentURL(applicationID, redirectURI, state string) (*url.URL, error) {
	consentURL, err := url.Parse("https://login.microsoftonline.com/common/adminconsent")
	if err != nil {
		return nil, err
	}

	query := url.Values{
		"client_id": []string{applicationID},
	}

	if redirectURI != "" {
		query.Set("redirect_uri", redirectURI)
	}

	if state != "" {
		query.Set("state", state)
	}

	consentURL.RawQuery = query.Encode()

	return consentURL, nil
}
