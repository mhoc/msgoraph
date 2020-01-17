package users

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

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
func (s *ServiceContext) CreateUser(createUser CreateUserRequest) (User, error) {
	body, err := client.GraphRequest(context.TODO(), s.client, http.MethodPost, "v1.0/users", nil, createUser)
	var data GetUserResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return User{}, err
	}
	return data.User, nil
}

// DeleteUser deletes an existing user by id or principal name.
func (s *ServiceContext) DeleteUser(userIDOrPrincipal string) error {
	reqURL := fmt.Sprintf("v1.0/users/%v", userIDOrPrincipal)
	_, err := client.GraphRequest(context.TODO(), s.client, http.MethodDelete, reqURL, nil, nil)
	return err
}

// GetUser returns a single user by id or principal name, with the Microsoft default fields
// provided, identical to those specified in UserDefaultFields.
func (s *ServiceContext) GetUser(userIDOrPrincipal string) (User, error) {
	return s.GetUserWithFields(userIDOrPrincipal, UserDefaultFields)
}

// GetUserWithFields returns a single user by id or principal name. You need to specify a list of
// fields you want to project on the user returned. You can specify UserDefaultFields or
// UserAllFields, or customize it depending on what you want.
func (s *ServiceContext) GetUserWithFields(userIDOrPrincipal string, projection []Field) (User, error) {
	if len(projection) == 0 {
		return User{}, fmt.Errorf("no fields provided in call to Users")
	}
	selectFields := ""
	for i, requestField := range projection {
		if i != 0 {
			selectFields += ","
		}
		selectFields += string(requestField)
	}
	v := url.Values{}
	v.Set("$select", selectFields)
	reqURL := fmt.Sprintf("v1.0/users/%v", userIDOrPrincipal)
	b, err := client.GraphRequest(context.TODO(), s.client, http.MethodGet, reqURL, v, nil)
	var data GetUserResponse
	err = json.Unmarshal(b, &data)
	if err != nil {
		return User{}, err
	}
	return data.User, nil
}

// ListUsers returns all users in the tenant, with each user projected with the Microsoft-defined
// default fields identical to UserDefaultFields.
func (s *ServiceContext) ListUsers() ([]User, error) {
	return s.ListUsersWithFields(UserDefaultFields)
}

// ListUsersWithFields returns the users on a tenant's azure instance. You need to specify a list of
// fields you want to project on the users returned. You can specify UserDefaultFields or
// UserAllFields, or customize it depending on what you want.
func (s *ServiceContext) ListUsersWithFields(projection []Field) ([]User, error) {
	getUserPage := func(url string) ([]User, string, error) {
		b, err := client.BasicGraphRequest(context.TODO(), s.client, http.MethodGet, url, nil)
		var data ListUsersResponse
		err = json.Unmarshal(b, &data)
		if err != nil {
			return nil, "", err
		}
		return data.Value, data.NextPage, nil
	}
	var users []User
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

// UpdateUser updates a user in the microsoft graph api, by userid or principal name, which is
// usually their email address. You can provide as few or many fields in the request as you'd like
// to update.
func (s *ServiceContext) UpdateUser(userIDOrPrincipal string, u UpdateUserRequest) error {
	reqURL := fmt.Sprintf("v1.0/users/%v", userIDOrPrincipal)
	_, err := client.GraphRequest(context.TODO(), s.client, http.MethodPatch, reqURL, nil, u)
	return err
}
