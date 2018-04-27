package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// Credentials stores enough information to authenticate a connection with some given
// TenantID. This includes the root clientID/clientSecret that authenticates a connection, as well
// as storage and logic to refresh access tokens when they expire.
type Credentials struct {
	AccessToken          string
	AccessTokenExpiresAt time.Time
	AccessTokenUpdating  sync.Mutex
	ClientID             string
	ClientSecret         string
	TenantID             string
}

// NewCredentials creates a new set of credentials. These credentials cannot immediately be used
// for fetching information from the Graph API; an access token first needs to be fetched. This
// will be done automatically on the first request to the Graph API, or manually by calling
// Credentials.RefreshAccessToken().
func NewCredentials(tenantID string, clientID string, clientSecret string) *Credentials {
	return &Credentials{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TenantID:     tenantID,
	}
}

// AuthEndpoint returns the oauth2 endpoint that is used to refresh the access token used to
// authenticate requests to the Graph API
func (c *Credentials) AuthEndpoint() string {
	return fmt.Sprintf("https://login.microsoftonline.com/%v/oauth2/v2.0/token", c.TenantID)
}

// RefreshAccessToken retrieves a 5 minute access token for the given set of credentials and upates
// it within the Credentials struct. This will refuse to update the token if the token hasn't
// yet expired, or if another goroutine is already doing the update.
func (c *Credentials) RefreshAccessToken() error {
	c.AccessTokenUpdating.Lock()
	defer c.AccessTokenUpdating.Unlock()
	if c.AccessToken != "" && c.AccessTokenExpiresAt.After(time.Now()) {
		return nil
	}
	resp, err := http.PostForm(c.AuthEndpoint(), url.Values{
		"client_id":     {c.ClientID},
		"client_secret": {c.ClientSecret},
		"grant_type":    {"client_credentials"},
		"scope":         {"https://graph.microsoft.com/.default"},
	})
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var data map[string]interface{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	serverErrCode, ok := data["error"].(string)
	if ok {
		serverErr, ok := data["error_description"].(string)
		if ok {
			return fmt.Errorf("%v: %v", serverErrCode, serverErr)
		}
		return fmt.Errorf(serverErrCode)
	}
	accessToken, ok := data["access_token"].(string)
	if !ok || accessToken == "" {
		return fmt.Errorf("no access token found in response")
	}
	durationSecs, ok := data["expires_in"].(float64)
	if !ok || durationSecs == 0 {
		return fmt.Errorf("no token duration found in response")
	}
	expiresAt := time.Now().Add(time.Duration(durationSecs) * time.Second)
	c.AccessToken = accessToken
	c.AccessTokenExpiresAt = expiresAt
	return nil
}
