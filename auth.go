package msgoraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// AccessToken exports information about an individual access token that is used by the
// tenant connection to determine when to update the token.
type AccessToken struct {
	ExpiresAt time.Time
	Token     string
}

// AuthEndpoint returns the oauth2 endpoint to which we should make a post request to retrieve
// a new oauth2 token.
func (t *Tenant) AuthEndpoint() string {
	return fmt.Sprintf("https://login.microsoftonline.com/%v/oauth2/v2.0/token", t.TenantID)
}

// GetAccessToken retrieves a 5 minute access token on the given tenant connection.
func (t *Tenant) GetAccessToken() (*AccessToken, error) {
	resp, err := http.PostForm(t.AuthEndpoint(), url.Values{
		"client_id":     {t.ClientID},
		"client_secret": {t.ClientSecret},
		"grant_type":    {"client_credentials"},
		"scope":         {"https://graph.microsoft.com/.default"},
	})
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}
	serverErrCode, ok := data["error"].(string)
	if ok {
		serverErr, ok := data["error_description"].(string)
		if ok {
			return nil, fmt.Errorf("%v: %v", serverErrCode, serverErr)
		}
		return nil, fmt.Errorf(serverErrCode)
	}
	accessToken, ok := data["access_token"].(string)
	if !ok || accessToken == "" {
		return nil, fmt.Errorf("no access token found in response")
	}
	durationSecs, ok := data["expires_in"].(float64)
	if !ok || durationSecs == 0 {
		return nil, fmt.Errorf("no token duration found in response")
	}
	expiresAt := time.Now().Add(time.Duration(durationSecs) * time.Second)
	return &AccessToken{
		ExpiresAt: expiresAt,
		Token:     accessToken,
	}, nil
}

// RefreshAccessToken can be called to force the client to refresh the access token for a given
// tenant. Generally consumers don't need to call this; it is all handled internally on every
// API request, but it is exposed in the event consumers find it necessary.
func (t *Tenant) RefreshAccessToken() error {
	token, err := t.GetAccessToken()
	if err != nil {
		return err
	}
	t.AccessToken = token
	return nil
}

// RefreshAccessTokenIfExpired checks the expiration on the current token and only refreshes it
// if it is expired.
func (t *Tenant) RefreshAccessTokenIfExpired() error {
	t.UpdatingAccessToken.Lock()
	defer t.UpdatingAccessToken.Unlock()
	if t.AccessToken.Token != "" && t.AccessToken.ExpiresAt.After(time.Now()) {
		return nil
	}
	return t.RefreshAccessToken()
}
