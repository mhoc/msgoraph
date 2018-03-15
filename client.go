// Package msgoraph implements client state, authentication, pagination, type helpers, and method
// helpers concerning accessing resources in the Microsoft Graph API.
package msgoraph

// Client maintains "connections" to multiple tenants' Azure instances. It takes care of auto-
// updating API tokens if they're expired, and maintaining separate tokens for each tenant
// we try to access. All operations on a Client are thread-safe.
type Client struct {
	ClientID     string
	ClientSecret string
	Cncts        map[string]*Tenant
}

// NewClient creates a new client connection to the azure graph api.
func NewClient(clientID, clientSecret string) *Client {
	return &Client{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Cncts:        make(map[string]*Tenant),
	}
}

// Tenant returns the connection for a specific tenant, by ID. If the tenant doesn't exist, it
// will be created, but no access token will be fetched.
func (c *Client) Tenant(tenantID string) *Tenant {
	cnct, in := c.Cncts[tenantID]
	if !in {
		c.Cncts[tenantID] = &Tenant{
			AccessToken:  &AccessToken{},
			ClientID:     c.ClientID,
			ClientSecret: c.ClientSecret,
			TenantID:     tenantID,
		}
	}
	cnct = c.Cncts[tenantID]
	return cnct
}
