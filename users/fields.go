package users

// Field can be provided to the user request functions to select which Fields
// are provided by Microsoft for each user. There's one for every root Field on the user object
// and they match up perfectly with the json names above. All of these have little comments
// on them for the purpose of supressing godoc linter warnings.
type Field string

const (
	// FieldID id
	FieldID Field = "id"
	// FieldAboutMe aboutMe
	FieldAboutMe Field = "aboutMe"
	// FieldAccountEnabled accountsEnabled
	FieldAccountEnabled Field = "accountEnabled"
	// FieldAssignedLicenses assignedLicenses
	FieldAssignedLicenses Field = "assignedLicenses"
	// FieldAssignedPlans assignedPlans
	FieldAssignedPlans Field = "assignedPlans"
	// FieldBirthday birthday
	FieldBirthday Field = "birthday"
	// FieldBusinessPhones businessPhones
	FieldBusinessPhones Field = "businessPhones"
	// FieldCity city
	FieldCity Field = "city"
	// FieldCompanyName companyName
	FieldCompanyName Field = "companyName"
	// FieldCountry country
	FieldCountry Field = "country"
	// FieldDepartment department
	FieldDepartment Field = "department"
	// FieldDisplayName displayName
	FieldDisplayName Field = "displayName"
	// FieldGivenName givenName
	FieldGivenName Field = "givenName"
	// FieldHireDate hireDate
	FieldHireDate Field = "hireDate"
	// FieldIMAddresses imAddresses
	FieldIMAddresses Field = "imAddresses"
	// FieldInterests interests
	FieldInterests Field = "interests"
	// FieldJobTitle jobTitle
	FieldJobTitle Field = "jobTitle"
	// FieldMail mail
	FieldMail Field = "mail"
	// FieldMailboxSettings mailboxSettings
	FieldMailboxSettings Field = "mailboxSettings"
	// FieldMailNickname mailNickname
	FieldMailNickname Field = "mailNickname"
	// FieldMobilePhone mobilePhone
	FieldMobilePhone Field = "mobilePhone"
	// FieldMySite mySite
	FieldMySite Field = "mySite"
	// FieldOfficeLocation officeLocation
	FieldOfficeLocation Field = "officeLocation"
	// FieldOnPremisesImmutableID onPremisesImmutableId
	FieldOnPremisesImmutableID Field = "onPremisesImmutableId"
	// FieldOnPremisesLastSyncDateTime onPremisesLastSyncDateTime
	FieldOnPremisesLastSyncDateTime Field = "onPremisesLastSyncDateTime"
	// FieldOnPremisesSecurityIdentifier onPremisesSecurityIdentifier
	FieldOnPremisesSecurityIdentifier Field = "onPremisesSecurityIdentifier"
	// FieldOnPremisesSyncEnabled onPremisesSyncEnabled
	FieldOnPremisesSyncEnabled Field = "onPremisesSyncEnabled"
	// FieldPasswordPolicies passwordPolicies
	FieldPasswordPolicies Field = "passwordPolicies"
	// FieldPasswordProfile passwordProfile
	FieldPasswordProfile Field = "passwordProfile"
	// FieldPastProjects pastProjects
	FieldPastProjects Field = "pastProjects"
	// FieldPostalCode postalCode
	FieldPostalCode Field = "postalCode"
	// FieldPreferredLanguage preferredLanguage
	FieldPreferredLanguage Field = "preferredLanguage"
	// FieldPreferredName preferredName
	FieldPreferredName Field = "preferredName"
	// FieldProvisionedPlans provisionedPlans
	FieldProvisionedPlans Field = "provisionedPlans"
	// FieldProxyAddresses proxyAddresses
	FieldProxyAddresses Field = "proxyAddresses"
	// FieldResponsibilities responsibilities
	FieldResponsibilities Field = "responsibilities"
	// FieldSchools schools
	FieldSchools Field = "schools"
	// FieldSkills skills
	FieldSkills Field = "skills"
	// FieldState state
	FieldState Field = "state"
	// FieldStreetAddress streetAddress
	FieldStreetAddress Field = "streetAddress"
	// FieldSurname surname
	FieldSurname Field = "surname"
	// FieldUsageLocation usageLocation
	FieldUsageLocation Field = "usageLocation"
	// FieldUserPrincipalName userPrincipalName
	FieldUserPrincipalName Field = "userPrincipalName"
	// FieldUserType userType
	FieldUserType Field = "userType"
)

var (
	// UserAllFields specifies every user field available for selection in api calls.
	UserAllFields = []Field{
		FieldID,
		FieldAboutMe,
		FieldAccountEnabled,
		FieldAssignedLicenses,
		FieldAssignedPlans,
		FieldBirthday,
		FieldBusinessPhones,
		FieldCity,
		FieldCompanyName,
		FieldCountry,
		FieldDepartment,
		FieldDisplayName,
		FieldGivenName,
		FieldHireDate,
		FieldIMAddresses,
		FieldInterests,
		FieldJobTitle,
		FieldMail,
		FieldMailboxSettings,
		FieldMailNickname,
		FieldMobilePhone,
		FieldMySite,
		FieldOfficeLocation,
		FieldOnPremisesImmutableID,
		FieldOnPremisesLastSyncDateTime,
		FieldOnPremisesSecurityIdentifier,
		FieldOnPremisesSyncEnabled,
		FieldPasswordPolicies,
		FieldPasswordProfile,
		FieldPastProjects,
		FieldPostalCode,
		FieldPreferredLanguage,
		FieldPreferredName,
		FieldProvisionedPlans,
		FieldProxyAddresses,
		FieldResponsibilities,
		FieldSchools,
		FieldSkills,
		FieldState,
		FieldStreetAddress,
		FieldSurname,
		FieldUsageLocation,
		FieldUserPrincipalName,
	}
	// UserDefaultFields specifies the Microsoft-specified default fields available for selection
	// in API calls.
	UserDefaultFields = []Field{
		FieldAccountEnabled,
		FieldBusinessPhones,
		FieldDisplayName,
		FieldGivenName,
		FieldID,
		FieldJobTitle,
		FieldMail,
		FieldMobilePhone,
		FieldOfficeLocation,
		FieldPreferredLanguage,
		FieldSurname,
		FieldUserPrincipalName,
	}
)
