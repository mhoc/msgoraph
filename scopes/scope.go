package scopes

import (
	"strings"
)

// PermissionType defines whether a permission is delegated or application.
type PermissionType string

// Scope contains information about the OAuth scopes exported by the Graph API. This includes the
// permission itself, as well as information about displaying it in a UI if you ever need to. This
// information is all pulled from the Graph API documentation at the time of this file's commit,
// and could change at any point (https://developer.microsoft.com/en-us/graph/docs/concepts/permissions_reference).
type Scope struct {
	AdminConsentRequired bool
	DisplayString        string
	Description          string
	Permission           string
	Type                 PermissionType
}

// Scopes is an alias to an array of scopes, with some additional logic on the type to make interfacing
// with the collection easier.
type Scopes []Scope

const (
	// PermissionTypeApplication application permissions
	PermissionTypeApplication = "application"
	// PermissionTypeAll both application and delegated permission types. This is mostly used
	// internally for querying and building lists on all the permissions in the graph api.
	PermissionTypeAll = "all"
	// PermissionTypeDelegated delegated permissions
	PermissionTypeDelegated = "delegated"
)

// All returns a list of every permission in the graph api by permission type.
func All(typ PermissionType) Scopes {
	var scopes Scopes
	if typ == PermissionTypeAll || typ == PermissionTypeApplication {
		scopes = append(scopes, []Scope{
			ApplicationContactsRead,
			ApplicationContactsReadWrite,
			ApplicationContactsRead,
			ApplicationContactsReadWrite,
			ApplicationDeviceReadWriteAll,
			ApplicationDirectoryReadAll,
			ApplicationDirectoryReadWriteAll,
			ApplicationPeopleReadAll,
			ApplicationUserReadAll,
			ApplicationUserReadWriteAll,
			ApplicationUserInviteAll,
			ApplicationUserExportAll,
		}...)
	}
	if typ == PermissionTypeAll || typ == PermissionTypeDelegated {
		scopes = append(scopes, []Scope{
			DelegatedCalendarsRead,
			DelegatedCalendarsReadShared,
			DelegatedCalendarsReadWrite,
			DelegatedCalendarsReadWriteShared,
			DelegatedContactsRead,
			DelegatedCalendarsRead,
			DelegatedCalendarsReadShared,
			DelegatedCalendarsReadWrite,
			DelegatedCalendarsReadWriteShared,
			DelegatedDeviceRead,
			DelegatedDeviceCommand,
			DelegatedDirectoryReadAll,
			DelegatedDirectoryReadWriteAll,
			DelegatedDirectoryAccessAsUser,
			DelegatedDeviceManagementAppsReadAll,
			DelegatedDeviceManagementAppsReadWriteAll,
			DelegatedDeviceManagementConfigurationReadAll,
			DelegatedDeviceManagementConfigurationReadWriteAll,
			DelegatedDeviceManagementManagedDevicesPrivilegedOperationsAll,
			DelegatedDeviceManagementManagedDevicesReadAll,
			DelegatedDeviceManagementManagedDevicesReadWriteAll,
			DelegatedDeviceManagementRBACReadAll,
			DelegatedDeviceManagementRBACReadWriteAll,
			DelegatedDeviceManagementServiceConfigReadAll,
			DelegatedDeviceManagementServiceConfigReadWriteAll,
			DelegatedEmail,
			DelegatedOfflineAccess,
			DelegatedOpenID,
			DelegatedProfile,
			DelegatedPeopleRead,
			DelegatedPeopleReadAll,
			DelegatedUserRead,
			DelegatedUserReadWrite,
			DelegatedUserReadBasicAll,
			DelegatedUserReadAll,
			DelegatedUserReadWriteAll,
			DelegatedUserInviteAll,
			DelegatedUserExportAll,
		}...)
	}
	return scopes
}

// Resolve takes a list of space separate scopes, also known as the output of QueryString, and
// reassembles them into a list of valid Scopes.
// You also need to specify whether you want these returned scopes to be delegated or application
// scopes, given that many permission names are shared between delegated and application
// permissions.
// This function will panic if PermissionTypeAll is provided.
func Resolve(scopeList string, preferType PermissionType) Scopes {
	if preferType == PermissionTypeAll {
		panic("must provide either PermissionTypeApplication or PermissionTypeDelegated")
	}
	allScopes := All(preferType)
	each := strings.Split(scopeList, " ")
	var foundScopes Scopes
	for _, s := range each {
		ss := allScopes.Find(s)
		if ss != nil {
			foundScopes = append(foundScopes, *ss)
		}
	}
	return foundScopes
}

// Find will locate the scope within the given list of scopes with the given permission name. To
// search across all available scopes, chain it with All().Find("user.read")
func (s Scopes) Find(permissionName string) *Scope {
	for _, ss := range s {
		if strings.ToLower(ss.Permission) == strings.ToLower(permissionName) {
			return &ss
		}
	}
	return nil
}

// HasScope returns true if the list of scopes contains the requested scope.
func (s Scopes) HasScope(scope Scope) bool {
	for _, iS := range s {
		if iS.Permission == scope.Permission && iS.Type == scope.Type {
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
