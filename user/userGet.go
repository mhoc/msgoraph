package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetUserResponse is the response to expect on a GetUser Request.
type GetUserResponse struct {
	Context string `json:"@odata.context"`
	User
}

// ListUsersResponse is the Response from the list users graph api endpoint
type ListUsersResponse struct {
	Context  string  `json:"@odata.context"`
	NextPage string  `json:"@odata.nextLink"`
	Value    []*User `json:"value"`
}

// GetUser returns a single user by id or principal name. You need to specify a list of fields you
// want to project on the user returned. You can specify msgraph.UserDefaultFields or
// msgraph.UserAllFields if you want, or customize it depending on what you want.
func GetUser(accessToken string, userIDOrPrincipal string, projection []Field) (*User, error) {
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
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", accessToken))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data GetUserResponse
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}
	return &data.User, nil
}

// ListUsers returns the users on a tenant's azure instance. You need to specify a list of fields you
// want to project on the users returned. You can specify msgraph.UserDefaultFields or
// msgraph.UserAllFields if you want, or customize it depending on what you want.
func ListUsers(accessToken string, projection []Field) ([]*User, error) {
	getUserPage := func(url string) ([]*User, string, error) {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, "", err
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", accessToken))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, "", err
		}
		b, err := ioutil.ReadAll(resp.Body)
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
