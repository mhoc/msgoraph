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

// Scopes is an alias to an array of scopes, with some additional logic on the type to make interfacing
// with the collection easier.
type Scopes []Scope

// HasScope returns true if the list of scopes contains the requested scope.
func (s Scopes) HasScope(scope Scope) bool {
	for _, iS := range s {
		if iS.Permission == scope.Permission && iS.Application == scope.Application && iS.Delegated == scope.Delegated {
			return true
		}
	}
	return false
}

// QueryString will turn a list of scopes into a query string suitable for consumption by
// the Graph API. Something like "offline_access user.read mail.read".
func (s Scopes) QueryString() string {
	qs := ""
	for _, scope := range s {
		qs += scope.Permission + " "
	}
	return qs
}
