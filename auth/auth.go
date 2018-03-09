package auth

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

// Endpoint returns the oauth2 endpoint to which we should make a post request to retrieve
// a new oauth2 token.
func Endpoint(tenantID string) string {
	return fmt.Sprintf("https://login.microsoftonline.com/%v/oauth2/v2.0/token", tenantID)
}

// GetAccessToken retrieves a 5 minute access token on the given tenant connection.
func GetAccessToken(clientID string, clientSecret string, tenantID string) (*AccessToken, error) {
	resp, err := http.PostForm(Endpoint(tenantID), url.Values{
		"client_id":     {clientID},
		"client_secret": {clientSecret},
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
