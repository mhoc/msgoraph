package msgoraph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// AccessToken contains information about an individual access token that is used by the tenant
// to determine when to update the token.
type AccessToken struct {
	ExpiresAt time.Time
	Token     string
}

// Tenant represents an individual connection to an azure tenant. Note that this copies the
// client id and client secret fields from the base client; this is primarily to make it nicer
// to write methods against this type, but in the future we could also use different clients ids
// and secrets for different tenants.
type Tenant struct {
	AccessToken  *AccessToken
	ClientID     string
	ClientSecret string
	TenantID     string
	// A mutex is used to protect the access token against refreshes by multiple threads. This is
	// all handled by the RefreshAccessTokenIfExpired function.
	UpdatingAccessToken sync.Mutex
}

// AuthEndpoint returns the oauth2 endpoint to which should be used when making a request to
// retrieve a new oauth2 token.
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

func (t *Tenant) request(method string, path string) ([]byte, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("https://graph.microsoft.com/v1.0/%v", path), nil)
	if err != nil {
		return nil, err
	}
	return t.requestCore(req)
}

func (t *Tenant) requestWithBody(method string, path string, body interface{}) ([]byte, error) {
	j, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, fmt.Sprintf("https://graph.microsoft.com/v1.0/%v", path), bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	return t.requestCore(req)
}

func (t *Tenant) requestWithParams(method string, path string, params url.Values) ([]byte, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("https://graph.microsoft.com/v1.0/%v?%v", path, params.Encode()), nil)
	if err != nil {
		return nil, err
	}
	return t.requestCore(req)
}

func (t *Tenant) requestCore(req *http.Request) ([]byte, error) {
	err := t.RefreshAccessTokenIfExpired()
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", t.AccessToken.Token))
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}
