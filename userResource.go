package msgoraph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// CreateUserRequest is all the available args you can set when creating a user.
type CreateUserRequest struct {
	AccountEnabled        bool                `json:"accountEnabled"`
	DisplayName           string              `json:"displayName"`
	OnPremisesImmutableID string              `json:"onPremisesImmutableId"`
	MailNickname          string              `json:"mailNickname"`
	PasswordProfile       UserPasswordProfile `json:"passwordProfile"`
	UserPrincipalName     string              `json:"userPrincipalName"`
}

// UserResource is a resource type for user information.
type UserResource struct {
	*Tenant
}

// Users create a UserResousrce from a tenant connection.
func (t *Tenant) Users() *UserResource {
	return &UserResource{
		Tenant: t,
	}
}

// Create creates a new user in the tenant.
func (ur *UserResource) Create(user CreateUserRequest) (*User, error) {
	err := ur.Tenant.RefreshAccessTokenIfExpired()
	if err != nil {
		return nil, err
	}
	j, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	reqURL := "https://graph.microsoft.com/v1.0/users"
	req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", ur.Tenant.AccessToken.Token))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data getUserResponse
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}
	return &data.User, nil
}

// Delete deletes an existing user by id or principal name.
func (ur *UserResource) Delete(userIDOrPrincipal string) error {
	err := ur.Tenant.RefreshAccessTokenIfExpired()
	if err != nil {
		return err
	}
	reqURL := fmt.Sprintf("https://graph.microsoft.com/v1.0/users/%v", userIDOrPrincipal)
	req, err := http.NewRequest("DELETE", reqURL, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", ur.Tenant.AccessToken.Token))
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}

// Get returns a single user by id or principal name, with the Microsoft default fields
// provided, identical to those specified in UserDefaultFields.
func (ur *UserResource) Get(userIDOrPrincipal string) (*User, error) {
	return ur.GetWithFields(userIDOrPrincipal, UserDefaultFields)
}

// GetWithFields returns a single user by id or principal name. You need to specify a list of
// fields you want to project on the user returned. You can specify UserDefaultFields or
// UserAllFields, or customize it depending on what you want.
func (ur *UserResource) GetWithFields(userIDOrPrincipal string, projection []UserField) (*User, error) {
	err := ur.Tenant.RefreshAccessTokenIfExpired()
	if err != nil {
		return nil, err
	}
	if len(projection) == 0 {
		return nil, fmt.Errorf("no fields provided in call to Users")
	}
	filter := "$select="
	for _, requestField := range projection {
		filter += string(requestField) + ","
	}
	reqURL := fmt.Sprintf("https://graph.microsoft.com/v1.0/users/%v?%v", userIDOrPrincipal, filter)
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", ur.Tenant.AccessToken.Token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data getUserResponse
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}
	return &data.User, nil
}

// List returns all users in the tenant, with each user projected with the Microsoft-defined
// default fields identical to UserAllFields.
func (ur *UserResource) List() ([]*User, error) {
	return ur.ListWithFields(UserAllFields)
}

// ListWithFields returns the users on a tenant's azure instance. You need to specify a list of
// fields you want to project on the users returned. You can specify UserDefaultFields or
// UserAllFields, or customize it depending on what you want.
func (ur *UserResource) ListWithFields(projection []UserField) ([]*User, error) {
	err := ur.Tenant.RefreshAccessTokenIfExpired()
	if err != nil {
		return nil, err
	}
	getUserPage := func(url string) ([]*User, string, error) {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, "", err
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", ur.Tenant.AccessToken.Token))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, "", err
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, "", err
		}
		var data listUsersResponse
		err = json.Unmarshal(b, &data)
		if err != nil {
			return nil, "", err
		}
		return data.Value, data.NextPage, nil
	}
	var users []*User
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
