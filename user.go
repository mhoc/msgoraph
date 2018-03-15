package msgoraph

const (
	// UserFieldID id
	UserFieldID UserField = "id"
	// UserFieldAboutMe aboutMe
	UserFieldAboutMe UserField = "aboutMe"
	// UserFieldAccountEnabled accountsEnabled
	UserFieldAccountEnabled UserField = "accountEnabled"
	// UserFieldAssignedLicenses assignedLicenses
	UserFieldAssignedLicenses UserField = "assignedLicenses"
	// UserFieldAssignedPlans assignedPlans
	UserFieldAssignedPlans UserField = "assignedPlans"
	// UserFieldBirthday birthday
	UserFieldBirthday UserField = "birthday"
	// UserFieldBusinessPhones businessPhones
	UserFieldBusinessPhones UserField = "businessPhones"
	// UserFieldCity city
	UserFieldCity UserField = "city"
	// UserFieldCompanyName companyName
	UserFieldCompanyName UserField = "companyName"
	// UserFieldCountry country
	UserFieldCountry UserField = "country"
	// UserFieldDepartment department
	UserFieldDepartment UserField = "department"
	// UserFieldDisplayName displayName
	UserFieldDisplayName UserField = "displayName"
	// UserFieldGivenName givenName
	UserFieldGivenName UserField = "givenName"
	// UserFieldHireDate hireDate
	UserFieldHireDate UserField = "hireDate"
	// UserFieldIMAddresses imAddresses
	UserFieldIMAddresses UserField = "imAddresses"
	// UserFieldInterests interests
	UserFieldInterests UserField = "interests"
	// UserFieldJobTitle jobTitle
	UserFieldJobTitle UserField = "jobTitle"
	// UserFieldMail mail
	UserFieldMail UserField = "mail"
	// UserFieldMailboxSettings mailboxSettings
	UserFieldMailboxSettings UserField = "mailboxSettings"
	// UserFieldMailNickname mailNickname
	UserFieldMailNickname UserField = "mailNickname"
	// UserFieldMobilePhone mobilePhone
	UserFieldMobilePhone UserField = "mobilePhone"
	// UserFieldMySite mySite
	UserFieldMySite UserField = "mySite"
	// UserFieldOfficeLocation officeLocation
	UserFieldOfficeLocation UserField = "officeLocation"
	// UserFieldOnPremisesImmutableID onPremisesImmutableId
	UserFieldOnPremisesImmutableID UserField = "onPremisesImmutableId"
	// UserFieldOnPremisesLastSyncDateTime onPremisesLastSyncDateTime
	UserFieldOnPremisesLastSyncDateTime UserField = "onPremisesLastSyncDateTime"
	// UserFieldOnPremisesSecurityIdentifier onPremisesSecurityIdentifier
	UserFieldOnPremisesSecurityIdentifier UserField = "onPremisesSecurityIdentifier"
	// UserFieldOnPremisesSyncEnabled onPremisesSyncEnabled
	UserFieldOnPremisesSyncEnabled UserField = "onPremisesSyncEnabled"
	// UserFieldPasswordPolicies passwordPolicies
	UserFieldPasswordPolicies UserField = "passwordPolicies"
	// UserFieldPasswordProfile passwordProfile
	UserFieldPasswordProfile UserField = "passwordProfile"
	// UserFieldPastProjects pastProjects
	UserFieldPastProjects UserField = "pastProjects"
	// UserFieldPostalCode postalCode
	UserFieldPostalCode UserField = "postalCode"
	// UserFieldPreferredLanguage preferredLanguage
	UserFieldPreferredLanguage UserField = "preferredLanguage"
	// UserFieldPreferredName preferredName
	UserFieldPreferredName UserField = "preferredName"
	// UserFieldProvisionedPlans provisionedPlans
	UserFieldProvisionedPlans UserField = "provisionedPlans"
	// UserFieldProxyAddresses proxyAddresses
	UserFieldProxyAddresses UserField = "proxyAddresses"
	// UserFieldResponsibilities responsibilities
	UserFieldResponsibilities UserField = "responsibilities"
	// UserFieldSchools schools
	UserFieldSchools UserField = "schools"
	// UserFieldSkills skills
	UserFieldSkills UserField = "skills"
	// UserFieldState state
	UserFieldState UserField = "state"
	// UserFieldStreetAddress streetAddress
	UserFieldStreetAddress UserField = "streetAddress"
	// UserFieldSurname surname
	UserFieldSurname UserField = "surname"
	// UserFieldUsageLocation usageLocation
	UserFieldUsageLocation UserField = "usageLocation"
	// UserFieldUserPrincipalName userPrincipalName
	UserFieldUserPrincipalName UserField = "userPrincipalName"
	// UserFieldUserType userType
	UserFieldUserType UserField = "userType"
)

var (
	// UserAllFields specifies every user field available for selection in api calls.
	UserAllFields = []UserField{
		UserFieldID,
		UserFieldAboutMe,
		UserFieldAccountEnabled,
		UserFieldAssignedLicenses,
		UserFieldAssignedPlans,
		UserFieldBirthday,
		UserFieldBusinessPhones,
		UserFieldCity,
		UserFieldCompanyName,
		UserFieldCountry,
		UserFieldDepartment,
		UserFieldDisplayName,
		UserFieldGivenName,
		UserFieldHireDate,
		UserFieldIMAddresses,
		UserFieldInterests,
		UserFieldJobTitle,
		UserFieldMail,
		UserFieldMailboxSettings,
		UserFieldMailNickname,
		UserFieldMobilePhone,
		UserFieldMySite,
		UserFieldOfficeLocation,
		UserFieldOnPremisesImmutableID,
		UserFieldOnPremisesLastSyncDateTime,
		UserFieldOnPremisesSecurityIdentifier,
		UserFieldOnPremisesSyncEnabled,
		UserFieldPasswordPolicies,
		UserFieldPasswordProfile,
		UserFieldPastProjects,
		UserFieldPostalCode,
		UserFieldPreferredLanguage,
		UserFieldPreferredName,
		UserFieldProvisionedPlans,
		UserFieldProxyAddresses,
		UserFieldResponsibilities,
		UserFieldSchools,
		UserFieldSkills,
		UserFieldState,
		UserFieldStreetAddress,
		UserFieldSurname,
		UserFieldUsageLocation,
		UserFieldUserPrincipalName,
	}
	// UserDefaultFields specifies the Microsoft-specified default fields available for selection
	// in API calls.
	UserDefaultFields = []UserField{
		UserFieldAccountEnabled,
		UserFieldBusinessPhones,
		UserFieldDisplayName,
		UserFieldGivenName,
		UserFieldID,
		UserFieldJobTitle,
		UserFieldMail,
		UserFieldMobilePhone,
		UserFieldOfficeLocation,
		UserFieldPreferredLanguage,
		UserFieldSurname,
		UserFieldUserPrincipalName,
	}
)

// getUserResponse is the response to expect on a GetUser Request.
type getUserResponse struct {
	Context string `json:"@odata.context"`
	User
}

// listUsersResponse is the Response from the list users graph api endpoint
type listUsersResponse struct {
	Context  string  `json:"@odata.context"`
	NextPage string  `json:"@odata.nextLink"`
	Value    []*User `json:"value"`
}

// User the user resource type in the microsoft graph qpi. Interpreted from this API
// documentation https://developer.microsoft.com/en-us/graph/docs/api-reference/v1.0/resources/user
type User struct {
	ID                           *string               `json:"id"`
	AboutMe                      *string               `json:"aboutMe"`
	AccountEnabled               *bool                 `json:"accountEnabled"`
	AssignedLicenses             []UserAssignedLicense `json:"assignedLicenses"`
	AssignedPlans                []AssignedPlan        `json:"assignedPlans"`
	Birthday                     *string               `json:"birthday"`
	BusinessPhones               []string              `json:"businessPhones"`
	CompanyName                  *string               `json:"companyName"`
	Country                      *string               `json:"country"`
	Department                   *string               `json:"department"`
	DisplayName                  *string               `json:"displayName"`
	GivenName                    *string               `json:"givenName"`
	HireDate                     *string               `json:"hireDate"`
	IMAddresses                  []string              `json:"imAddresses"`
	Interests                    []string              `json:"interests"`
	JobTitle                     *string               `json:"jobTitle"`
	Mail                         *string               `json:"mail"`
	MailboxSettings              *UserMailboxSettings  `json:"mailboxSettings"`
	MailNickname                 *string               `json:"mailNickname"`
	MobilePhone                  *string               `json:"mobilePhone"`
	MySite                       *string               `json:"mySite"`
	OfficeLocation               *string               `json:"officeLocation"`
	OnPremisesImmutableID        *string               `json:"onPremisesImmutableId"`
	OnPremisesLastSyncDateTime   *string               `json:"onPremisesLastSyncDateTime"`
	OnPremisesSecurityIdentifier *string               `json:"onPremisesSecurityIdentifier"`
	OnPremisesSyncEnabled        *bool                 `json:"onPremisesSyncEnabled"`
	PasswordPolicies             *string               `json:"passwordPolicies"`
	PasswordProfile              *UserPasswordProfile  `json:"passwordProfile"`
	PastProjects                 []string              `json:"pastProjects"`
	PostalCode                   *string               `json:"postalCode"`
	PreferredLanguage            *string               `json:"preferredLanguage"`
	PreferredName                *string               `json:"preferredName"`
	ProvisionedPlans             []ProvisionedPlan     `json:"provisionedPlans"`
	ProxyAddresses               []string              `json:"proxyAddresses"`
	Responsibilities             *string               `json:"responsibilities"`
	Schools                      *string               `json:"schools"`
	Skills                       *string               `json:"skills"`
	State                        *string               `json:"state"`
	StreetAddress                *string               `json:"streetAddress"`
	Surname                      *string               `json:"surname"`
	UsageLocation                *string               `json:"string"`
	UserPrincipalName            *string               `json:"userPrincipalName"`
	UserType                     *string               `json:"userType"`
}

// UserAssignedLicense represents a license assigned to a user
type UserAssignedLicense struct {
	DisabledPlans []string `json:"disabledPlans"`
	SKUID         *string  `json:"skuId"`
}

// UserAutomaticRepliesSetting configuration settings to automatically notify the sender of an
// incoming email with a message from the signed-in user. For example, an automatic reply to
// notify that the signed-in user is unavailable to respond to emails.
type UserAutomaticRepliesSetting struct {
	ExternalAudience       string `json:"externalAudience"`
	ExternalReplyMessage   string `json:"externalReplyMessage"`
	InternalReplyMessage   string `json:"InternalReplyMessage"`
	ScheduledEndDateTime   string `json:"scheduledEndDateTime"`
	ScheduledStartDateTime string `json:"scheduledStartDateTime"`
	Status                 string `json:"status"`
}

// UserField can be provided to the user request functions to select which UserFields
// are provided by Microsoft for each user. There's one for every root UserField on the user object
// and they match up perfectly with the json names above. All of these have little comments
// on them for the purpose of supressing godoc linter warnings.
type UserField string

// UserLocaleInfo contains information about the locale, including the preferred language and
// country/region, of the signed-in user.
type UserLocaleInfo struct {
	DisplayName string `json:"displayName"`
	Locale      string `json:"locale"`
}

// UserMailboxSettings Settings for the primary mailbox of the signed-in user.
type UserMailboxSettings struct {
	AutomaticRepliesSetting UserAutomaticRepliesSetting `json:"automaticRepliesSetting"`
	Language                UserLocaleInfo              `json:"localeInfo"`
	TimeZone                string                      `json:"timeZone"`
}

// UserPasswordProfile Contains the password profile associated with a user.
type UserPasswordProfile struct {
	ForceChangePasswordNextSignIn bool   `json:"forceChangePasswordNextSignIn"`
	Password                      string `json:"password"`
}
