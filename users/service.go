package users

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/mhoc/msgoraph/client"
)

// CreateUserRequest is all the available args you can set when creating a user.
type CreateUserRequest struct {
	AccountEnabled        bool            `json:"accountEnabled"`
	DisplayName           string          `json:"displayName"`
	OnPremisesImmutableID string          `json:"onPremisesImmutableId"`
	MailNickname          string          `json:"mailNickname"`
	PasswordProfile       PasswordProfile `json:"passwordProfile"`
	UserPrincipalName     string          `json:"userPrincipalName"`
}

// GetUserResponse is the response to expect on a GetUser Request.
type GetUserResponse struct {
	Context string `json:"@odata.context"`
	User
}

// ListUsersResponse is the Response from the list users graph api endpoint
type ListUsersResponse struct {
	Context  string `json:"@odata.context"`
	NextPage string `json:"@odata.nextLink"`
	Value    []User `json:"value"`
}

// ServiceContext represents a namespace under which all of the operations against user-namespaced
// resources are accessed.
type ServiceContext struct {
	client client.Client
}

// Service creates a new users.ServiceContext with the given authentication credentials.
func Service(client client.Client) *ServiceContext {
	return &ServiceContext{client: client}
}

// UpdateUserRequest contains the request body to update a user.
type UpdateUserRequest struct {
	AboutMe               string            `json:"aboutMe"`
	AccountEnabled        string            `json:"accountEnabled"`
	AssignedLicenses      []AssignedLicense `json:"assignedLicenses"`
	Birthday              string            `json:"birthday"`
	City                  string            `json:"city"`
	Country               string            `json:"country"`
	Department            string            `json:"department"`
	DisplayName           string            `json:"displayName"`
	GivenName             string            `json:"givenName"`
	HireDate              string            `json:"hireDate"`
	Interests             []string          `json:"interests"`
	JobTitle              string            `json:"jobTitle"`
	MailNickname          string            `json:"mailNickname"`
	MobilePhone           string            `json:"mobilePhone"`
	MySite                string            `json:"mySite"`
	OfficeLocation        string            `json:"officeLocation"`
	OnPremisesImmutableID string            `json:"onPremisesImmutableId"`
	PasswordPolicies      string            `json:"passwordPolicies"`
	PasswordProfile       PasswordProfile   `json:"passwordProfile"`
	PastProjects          []string          `json:"pastProjects"`
	PostalCode            string            `json:"postalCode"`
	PreferredLanguage     string            `json:"preferredLanguage"`
	PreferredName         string            `json:"preferredName"`
	Responsibilities      []string          `json:"responsibilities"`
	Schools               []string          `json:"schools"`
	Skills                []string          `json:"skills"`
	State                 string            `json:"state"`
	StreetAddress         string            `json:"streetAddress"`
	Surname               string            `json:"surname"`
	UsageLocation         string            `json:"usageLocation"`
	UserPrincipalName     string            `json:"userPrincipalName"`
	UserType              string            `json:"userType"`
}

// CreateUser creates a new user in the tenant.
func (s *ServiceContext) CreateUser(ctx context.Context, createUser CreateUserRequest) (User, error) {
	body, err := client.GraphRequest(ctx, s.client, http.MethodPost, "/v1.0/users", nil, createUser)
	var data GetUserResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return User{}, err
	}
	return data.User, nil
}

// DeleteUser deletes an existing user by id or principal name.
func (s *ServiceContext) DeleteUser(ctx context.Context, userIDOrPrincipal string) error {
	reqURL := "/v1.0/users/" + url.PathEscape(userIDOrPrincipal)
	_, err := client.GraphRequest(ctx, s.client, http.MethodDelete, reqURL, nil, nil)
	return err
}

// GetUser returns a single user by id or principal name, with the Microsoft default fields
// provided, identical to those specified in UserDefaultFields.
func (s *ServiceContext) GetUser(ctx context.Context, userIDOrPrincipal string) (User, error) {
	return s.GetUserWithFields(ctx, userIDOrPrincipal, UserDefaultFields)
}

// GetUserWithFields returns a single user by id or principal name. You need to specify a list of
// fields you want to project on the user returned. You can specify UserDefaultFields or
// UserAllFields, or customize it depending on what you want.
func (s *ServiceContext) GetUserWithFields(ctx context.Context, userIDOrPrincipal string, projection []Field) (User, error) {
	query := s.selectFields(url.Values{}, projection)
	reqURL := "/v1.0/users/" + url.PathEscape(userIDOrPrincipal)
	b, err := client.GraphRequest(ctx, s.client, http.MethodGet, reqURL, query, nil)
	var data GetUserResponse
	err = json.Unmarshal(b, &data)
	if err != nil {
		return User{}, err
	}
	return data.User, nil
}

func (s *ServiceContext) selectFields(query url.Values, fields []Field) url.Values {
	if len(fields) == 0 {
		return query
	}
	selectFields := strings.Builder{}
	for i, requestField := range fields {
		if i != 0 {
			selectFields.WriteByte(',')
		}
		selectFields.WriteString(string(requestField))
	}
	query.Set("$select", selectFields.String())
	return query
}

// ListUsers returns all users in the tenant, with each user projected with the default fields
// returned by the API.
func (s *ServiceContext) ListUsers(ctx context.Context) ([]User, error) {
	return s.ListUsersWithFields(ctx, nil)
}

// ListUsersWithFields returns the users on a tenant's Azure instance. The provided fields are
// used as the `$select` parameter to the API. You can specify `UserDefaultFields` or
// `UserAllFields`, or customize it depending on what you want. If you specify no fields, the
// API will returns its default fields.
func (s *ServiceContext) ListUsersWithFields(ctx context.Context, projection []Field) ([]User, error) {
	query := s.selectFields(url.Values{}, projection)
	usersURL, err := client.GraphAPIRootURL.Parse("/v1.0/users")
	if err != nil {
		return nil, err
	}
	usersURL.RawQuery = query.Encode()
	nextURL := usersURL.String()
	var users []User
	for nextURL != "" {
		pageUsers, next, err := s.getUserPage(ctx, nextURL)
		if err != nil {
			return nil, err
		}
		for i := range pageUsers {
			users = append(users, pageUsers[i])
		}
		nextURL = next
	}
	return users, nil
}

func (s *ServiceContext) getUserPage(ctx context.Context, url string) ([]User, string, error) {
	b, err := client.BasicGraphRequest(ctx, s.client, http.MethodGet, url, nil)
	if err != nil {
		return nil, "", err
	}
	var data ListUsersResponse
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, "", err
	}
	return data.Value, data.NextPage, nil
}

// UpdateUser updates a user in the microsoft graph api, by userid or principal name, which is
// usually their email address. You can provide as few or many fields in the request as you'd like
// to update.
func (s *ServiceContext) UpdateUser(ctx context.Context, userIDOrPrincipal string, u UpdateUserRequest) error {
	reqURL := "/v1.0/users/" + url.PathEscape(userIDOrPrincipal)
	_, err := client.GraphRequest(ctx, s.client, http.MethodPatch, reqURL, nil, u)
	return err
}
