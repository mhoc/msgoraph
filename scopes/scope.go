package scopes

// Scope contains information about the OAuth scopes exported by the Graph API. This includes the
// permission itself, as well as information about displaying it in a UI if you ever need to. This
// information is all pulled from the Graph API documentation at the time of this file's commit,
// and could change at any point (https://developer.microsoft.com/en-us/graph/docs/concepts/permissions_reference).
type Scope struct {
	AdminConsentRequired bool
	Application          bool
	Delegated            bool
	DisplayString        string
	Description          string
	Permission           string
}

// CreateQueryString will turn a list of scopes into a query string suitable for consumption by
// the Graph API. Something like "offline_access user.read mail.read".
func CreateQueryString(scopes []Scope) string {
	qs := ""
	for _, scope := range scopes {
		qs += scope.Permission + " "
	}
	return qs
}
