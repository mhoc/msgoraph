package user

import (
	"github.com/mhoc/msgoraph/plan"
)

// User the user resource type in the microsoft graph qpi. Interpreted from this API
// documentation https://developer.microsoft.com/en-us/graph/docs/api-reference/v1.0/resources/user
type User struct {
	ID                           *string                `json:"id"`
	AboutMe                      *string                `json:"aboutMe"`
	AccountEnabled               *bool                  `json:"accountEnabled"`
	AssignedLicenses             []AssignedLicense      `json:"assignedLicenses"`
	AssignedPlans                []plan.AssignedPlan    `json:"assignedPlans"`
	Birthday                     *string                `json:"birthday"`
	BusinessPhones               []string               `json:"businessPhones"`
	CompanyName                  *string                `json:"companyName"`
	Country                      *string                `json:"country"`
	Department                   *string                `json:"department"`
	DisplayName                  *string                `json:"displayName"`
	GivenName                    *string                `json:"givenName"`
	HireDate                     *string                `json:"hireDate"`
	IMAddresses                  []string               `json:"imAddresses"`
	Interests                    []string               `json:"interests"`
	JobTitle                     *string                `json:"jobTitle"`
	Mail                         *string                `json:"mail"`
	MailboxSettings              *MailboxSettings       `json:"mailboxSettings"`
	MailNickname                 *string                `json:"mailNickname"`
	MobilePhone                  *string                `json:"mobilePhone"`
	MySite                       *string                `json:"mySite"`
	OfficeLocation               *string                `json:"officeLocation"`
	OnPremisesImmutableID        *string                `json:"onPremisesImmutableId"`
	OnPremisesLastSyncDateTime   *string                `json:"onPremisesLastSyncDateTime"`
	OnPremisesSecurityIdentifier *string                `json:"onPremisesSecurityIdentifier"`
	OnPremisesSyncEnabled        *bool                  `json:"onPremisesSyncEnabled"`
	PasswordPolicies             *string                `json:"passwordPolicies"`
	PasswordProfile              *PasswordProfile       `json:"passwordProfile"`
	PastProjects                 []string               `json:"pastProjects"`
	PostalCode                   *string                `json:"postalCode"`
	PreferredLanguage            *string                `json:"preferredLanguage"`
	PreferredName                *string                `json:"preferredName"`
	ProvisionedPlans             []plan.ProvisionedPlan `json:"provisionedPlans"`
	ProxyAddresses               []string               `json:"proxyAddresses"`
	Responsibilities             *string                `json:"responsibilities"`
	Schools                      *string                `json:"schools"`
	Skills                       *string                `json:"skills"`
	State                        *string                `json:"state"`
	StreetAddress                *string                `json:"streetAddress"`
	Surname                      *string                `json:"surname"`
	UsageLocation                *string                `json:"string"`
	UserPrincipalName            *string                `json:"userPrincipalName"`
	UserType                     *string                `json:"userType"`
}

// StringShort returns a single line short string representation of the user, with tabs
// separating each field so it is appropriate for usage with a tabwriter.
func (u User) StringShort() string {
	s := *u.UserPrincipalName
	s += "\t"
	if u.GivenName != nil {
		s += *u.GivenName
	}
	s += "\t"
	if u.Surname != nil {
		s += *u.Surname
	}
	return s
}
