package user

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
