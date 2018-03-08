package msgoraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/mhoc/msgoraph/user"
)

// Client maintains "connections" to multiple tenants' Azure instances. It takes care of auto-
// updating API tokens if they're expired, and maintaining separate tokens for each tenant
// we try to access. All operations on a Client are thread-safe.
type Client struct {
	ClientID     string
	ClientSecret string
	Cncts        map[string]*TenantCnct
}

// TenantCnct represents an individual connection to an azure tenant. Note that this copies the
// client id and client secret fields from the base client; this is primarily to make it nicer
// to write methods against this type, but in the future we could also use different clients ids
// and secrets for different tenants.
type TenantCnct struct {
	AccessToken          string
	AccessTokenExpiresAt time.Time
	ClientID             string
	ClientSecret         string
	TenantID             string
	// A mutex is used to protect the access token against refreshes by multiple threads. This is
	// all handled by the RefreshAccessTokenIfExpired function.
	UpdatingAccessToken sync.Mutex
}

// NewClient creates a new client connection to the azure graph api.
func NewClient(clientID, clientSecret string) *Client {
	return &Client{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Cncts:        make(map[string]*TenantCnct),
	}
}

// Tenant returns the connection for a specific tenant, by ID. If the tenant doesn't exist, it
// will be created, but no access token will be fetched.
func (c *Client) Tenant(tenantID string) *TenantCnct {
	cnct, in := c.Cncts[tenantID]
	if !in {
		c.Cncts[tenantID] = &TenantCnct{
			ClientID:     c.ClientID,
			ClientSecret: c.ClientSecret,
			TenantID:     tenantID,
		}
	}
	cnct = c.Cncts[tenantID]
	return cnct
}

// AuthEndpoint returns the oauth2 endpoint to which we should make a post request to retrieve
// a new oauth2 token.
func (t *TenantCnct) AuthEndpoint() string {
	return fmt.Sprintf("https://login.microsoftonline.com/%v/oauth2/v2.0/token", t.TenantID)
}

// RefreshAccessTokenIfExpired refreshes the 5 minute access token on the given tenant connection.
// This only actually makes a request to Microsoft if the stored token is expired.
func (t *TenantCnct) RefreshAccessTokenIfExpired() error {
	t.UpdatingAccessToken.Lock()
	defer t.UpdatingAccessToken.Unlock()
	if t.AccessToken != "" && t.AccessTokenExpiresAt.After(time.Now()) {
		return nil
	}
	authURL := t.AuthEndpoint()
	resp, err := http.PostForm(authURL, url.Values{
		"client_id":     {t.ClientID},
		"client_secret": {t.ClientSecret},
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
	t.AccessToken = accessToken
	expiresIn := time.Duration(data["expires_in"].(float64))
	t.AccessTokenExpiresAt = time.Now().Add(expiresIn * time.Second)
	return nil
}

// User returns a single user by id or principal name. You need to specify a list of fields you
// want to project on the user returned. You can specify msgraph.UserDefaultFields or
// msgraph.UserAllFields if you want, or customize it depending on what you want.
func (t *TenantCnct) User(userIDOrPrincipal string, projection []user.Field) (user.User, error) {
	var u user.User
	err := t.RefreshAccessTokenIfExpired()
	if err != nil {
		return u, err
	}
	if len(projection) == 0 {
		return u, fmt.Errorf("no fields provided in call to Users")
	}
	filter := "$select="
	for _, requestField := range projection {
		filter += string(requestField) + ","
	}
	reqURL := fmt.Sprintf("https://graph.microsoft.com/v1.0/users/%v?%v", userIDOrPrincipal, filter)
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return u, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", t.AccessToken))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return u, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return u, err
	}
	var data user.GetUserResponse
	err = json.Unmarshal(b, &data)
	if err != nil {
		return u, err
	}
	return data.User, nil
}

// Users returns the users on a tenant's azure instance. You need to specify a list of fields you
// want to project on the users returned. You can specify msgraph.UserDefaultFields or
// msgraph.UserAllFields if you want, or customize it depending on what you want.
func (t *TenantCnct) Users(projection []user.Field) ([]user.User, error) {
	err := t.RefreshAccessTokenIfExpired()
	if err != nil {
		return nil, err
	}
	getUserPage := func(url string) ([]user.User, string, error) {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, "", err
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", t.AccessToken))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, "", err
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, "", err
		}
		var data user.ListUsersResponse
		err = json.Unmarshal(b, &data)
		if err != nil {
			return nil, "", err
		}
		return data.Value, data.NextPage, nil
	}
	var users []user.User
	if len(projection) == 0 {
		return nil, fmt.Errorf("no fields provided in call to Users")
	}
	filter := "$select="
	for _, requestField := range projection {
		filter += string(requestField) + ","
	}
	nextURL := fmt.Sprintf("https://graph.microsoft.com/v1.0/users?%v", filter)
	for nextURL != "" {
		pageUsers, next, err := getUserPage(nextURL)
		if err != nil {
			return nil, err
		}
		for _, pageUser := range pageUsers {
			users = append(users, pageUser)
		}
		nextURL = next
	}
	return users, nil
}
