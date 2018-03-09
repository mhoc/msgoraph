package msgoraph

import (
	"sync"
	"time"

	"github.com/mhoc/msgoraph/auth"
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

// RefreshAccessToken can be called to force the client to refresh the access token for a given
// tenant. Generally consumers don't need to call this; it is all handled internally on every
// API request, but it is exposed in the event consumers find it necessary.
func (t *TenantCnct) RefreshAccessToken() error {
	token, err := auth.GetAccessToken(t.ClientID, t.ClientSecret, t.TenantID)
	if err != nil {
		return err
	}
	t.AccessToken = token.Token
	t.AccessTokenExpiresAt = token.ExpiresAt
	return nil
}

// RefreshAccessTokenIfExpired checks the expiration on the current token and only refreshes it
// if it is expired.
func (t *TenantCnct) RefreshAccessTokenIfExpired() error {
	t.UpdatingAccessToken.Lock()
	defer t.UpdatingAccessToken.Unlock()
	if t.AccessToken != "" && t.AccessTokenExpiresAt.After(time.Now()) {
		return nil
	}
	return t.RefreshAccessToken()
}

// User returns a single user from azure by user id or principal name (usually their email address),
// projecting the user with the default fields returned by Microsoft.
func (t *TenantCnct) User(userIDOrPrincipal string) (*user.User, error) {
	err := t.RefreshAccessTokenIfExpired()
	if err != nil {
		return nil, err
	}
	return user.GetUser(t.AccessToken, userIDOrPrincipal, user.DefaultFields)
}

// UserWithFields returns a single user from azure by user id or principal name (usually their email
// address), and allows you to specify the projection of fields on the returned user.
func (t *TenantCnct) UserWithFields(userIDOrPrincipal string, projection []user.Field) (*user.User, error) {
	err := t.RefreshAccessTokenIfExpired()
	if err != nil {
		return nil, err
	}
	return user.GetUser(t.AccessToken, userIDOrPrincipal, projection)
}

// Users returns all of the users on the given tenant, with each user projected with the default
// fields provided by Microsoft.
func (t *TenantCnct) Users() ([]*user.User, error) {
	err := t.RefreshAccessTokenIfExpired()
	if err != nil {
		return nil, err
	}
	return user.ListUsers(t.AccessToken, user.DefaultFields)
}

// UsersWithFields returns all of the users on the given tenant, with each user projected with the
// projection you provide.
func (t *TenantCnct) UsersWithFields(projection []user.Field) ([]*user.User, error) {
	err := t.RefreshAccessTokenIfExpired()
	if err != nil {
		return nil, err
	}
	return user.ListUsers(t.AccessToken, projection)
}
