package scopes

var (
	// ApplicationPeopleReadAll Read all users' relevant people lists
	ApplicationPeopleReadAll = Scope{
		AdminConsentRequired: true,
		Description:          "Allows the app to read a scored list of people relevant to the signed-in user or other users in the signed-in user's organization. The list can include local contacts, contacts from social networking or your organization's directory, and people from recent communications (such as email and Skype). Also allows the app to search the entire directory of the signed-in user's organization.",
		DisplayString:        "Read all users' relevant people lists",
		Permission:           "People.Read.All",
		Type:                 PermissionTypeApplication,
	}
	// DelegatedPeopleRead Read users' relevant people lists
	DelegatedPeopleRead = Scope{
		Description:   "Allows the app to read a scored list of people relevant to the signed-in user. The list can include local contacts, contacts from social networking or your organization's directory, and people from recent communications (such as email and Skype).",
		DisplayString: "Read users' relevant people lists",
		Permission:    "People.Read",
		Type:          PermissionTypeDelegated,
	}
	// DelegatedPeopleReadAll Read all users' relevant people lists
	DelegatedPeopleReadAll = Scope{
		AdminConsentRequired: true,
		Description:          "Allows the app to read a scored list of people relevant to the signed-in user or other users in the signed-in user's organization. The list can include local contacts, contacts from social networking or your organization's directory, and people from recent communications (such as email and Skype). Also allows the app to search the entire directory of the signed-in user's organization.",
		DisplayString:        "Read all users' relevant people lists",
		Permission:           "People.Read.All",
		Type:                 PermissionTypeDelegated,
	}
)
