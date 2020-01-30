package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/mhoc/msgoraph/scopes"
)

var _ Client = (*Headless)(nil)

// Headless is used to authenticate requests in the context of a backend app. This is the most
// common way for applications to authenticate with the api.
type Headless struct {
	ApplicationID      string
	ApplicationSecret  string
	Tenant             string
	RefreshToken       string
	RequestCredentials *RequestCredentials
	Scopes             scopes.Scopes
	Client             *http.Client
}

// NewHeadless creates a new headless connection.
func NewHeadless(applicationID string, applicationSecret string, scopes scopes.Scopes) *Headless {
	return &Headless{
		ApplicationID:      applicationID,
		ApplicationSecret:  applicationSecret,
		Tenant:             "common",
		RequestCredentials: &RequestCredentials{},
		Scopes:             scopes,
		Client:             http.DefaultClient,
	}
}

// HTTPClient returns the `*http.Client` to use for this `Headless` instance.
func (h Headless) HTTPClient() *http.Client {
	return h.Client
}

// Credentials returns back the set of credentials used for every request.
func (h Headless) Credentials() *RequestCredentials {
	return h.RequestCredentials
}

// InitializeCredentials will make an initial oauth2 token request for a new token.
func (h Headless) InitializeCredentials(ctx context.Context) error {
	h.RequestCredentials.AccessTokenUpdating.Lock()
	defer h.RequestCredentials.AccessTokenUpdating.Unlock()
	if h.RequestCredentials.AccessToken != "" && h.RequestCredentials.AccessTokenExpiresAt.After(time.Now()) {
		return nil
	}
	tokenURI, err := url.Parse("https://login.microsoftonline.com/" + url.PathEscape(h.Tenant) + "/oauth2/v2.0/token")
	if err != nil {
		return err
	}
	resp, err := postForm(
		ctx,
		h.HTTPClient(),
		tokenURI.String(),
		url.Values{
			"client_id":     {h.ApplicationID},
			"client_secret": {h.ApplicationSecret},
			"grant_type":    {"client_credentials"},
			"scope":         {"https://graph.microsoft.com/.default"},
		},
	)
	if err != nil {
		return err
	}
	b, err := handleResp(ctx, resp)
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
	if h.Scopes.HasScope(scopes.DelegatedOfflineAccess) {
		refreshToken, ok := data["refresh_token"].(string)
		if !ok || refreshToken == "" {
			return fmt.Errorf("no refresh token found in response")
		}
		h.RefreshToken = refreshToken
	}
	expiresAt := time.Now().Add(time.Duration(durationSecs) * time.Second)
	h.RequestCredentials.AccessToken = accessToken
	h.RequestCredentials.AccessTokenExpiresAt = expiresAt
	return nil
}

// RefreshCredentials will refresh the connection credentials. This just proxies through to
// InitializeCredentials, because in the context of a headless appliction we should probably already
// have the application secret key.
func (h Headless) RefreshCredentials(ctx context.Context) error {
	return h.InitializeCredentials(ctx)
}

// ConsentURL builds the consent URL needed to add this application to a target Azure domain.
func (h Headless) ConsentURL(redirectURL, state string) (*url.URL, error) {
	return ConsentURL(h.ApplicationID, redirectURL, state)
}
