package scopes

var (
	// ApplicationUserReadAll Read all users' full profiles
	ApplicationUserReadAll = Scope{
		AdminConsentRequired: true,
		Application:          true,
		Description:          "Allows the app to read the full set of profile properties, group membership, reports and managers of other users in your organization, without a signed-in user.",
		DisplayString:        "Read all users' full profiles",
		Permission:           "User.Read.All",
	}
	// ApplicationUserReadWriteAll Read and write all users' full profiles
	ApplicationUserReadWriteAll = Scope{
		AdminConsentRequired: true,
		Application:          true,
		Description:          "Allows the app to read and write the full set of profile properties, group membership, reports and managers of other users in your organization, without a signed-in user.  Also allows the app to create and delete non-administrative users. Does not allow reset of user passwords.",
		DisplayString:        "Read and write all users' full profiles",
		Permission:           "User.ReadWrite.All",
	}
	// ApplicationUserInviteAll Invite guest users to the organization
	ApplicationUserInviteAll = Scope{
		AdminConsentRequired: true,
		Application:          true,
		Description:          "Allows the app to invite guest users to your organization, without a signed-in user.",
		DisplayString:        "Invite guest users to the organization",
		Permission:           "User.Invite.All",
	}
	// ApplicationUserExportAll Export users' data
	ApplicationUserExportAll = Scope{
		AdminConsentRequired: true,
		Application:          true,
		Description:          "Allows the app to export organizational users' data, without a signed-in user.",
		DisplayString:        "Export users' data",
		Permission:           "User.Export.All",
	}
	// DelegatedUserRead Sign-in and read user profile
	DelegatedUserRead = Scope{
		Delegated:     true,
		Description:   "Allows users to sign-in to the app, and allows the app to read the profile of signed-in users. It also allows the app to read basic company information of signed-in users.",
		DisplayString: "Sign-in and read user profile",
		Permission:    "User.Read",
	}
	// DelegatedUserReadWrite Read and write access to user profile
	DelegatedUserReadWrite = Scope{
		Delegated:     true,
		Description:   "Allows the app to read your profile. It also allows the app to update your profile information on your behalf.",
		DisplayString: "Read and write access to user profile",
		Permission:    "User.ReadWrite",
	}
	// DelegatedUserReadBasicAll Read all users' basic profiles
	DelegatedUserReadBasicAll = Scope{
		Delegated:     true,
		Description:   "Allows the app to read a basic set of profile properties of other users in your organization on behalf of the signed-in user. This includes display name, first and last name, email address and photo.",
		DisplayString: "Read all users' basic profiles",
		Permission:    "User.ReadBasic.All",
	}
	// DelegatedUserReadAll Read all users' full profiles
	DelegatedUserReadAll = Scope{
		AdminConsentRequired: true,
		Delegated:            true,
		Description:          "Allows the app to read the full set of profile properties, reports, and managers of other users in your organization, on behalf of the signed-in user.",
		DisplayString:        "Read all users' full profiles",
		Permission:           "User.Read.All",
	}
	// DelegatedUserReadWriteAll Read and write all users' full profiles
	DelegatedUserReadWriteAll = Scope{
		AdminConsentRequired: true,
		Delegated:            true,
		Description:          "Allows the app to read and write the full set of profile properties, reports, and managers of other users in your organization, on behalf of the signed-in user. Also allows the app to create and delete users as well as reset user passwords on behalf of the signed-in user.",
		DisplayString:        "Read and write all users' full profiles",
		Permission:           "User.ReadWrite.All",
	}
	// DelegatedUserInviteAll Invite guest users to the organization
	DelegatedUserInviteAll = Scope{
		AdminConsentRequired: true,
		Delegated:            true,
		Description:          "Allows the app to invite guest users to your organization, on behalf of the signed-in user.",
		DisplayString:        "Invite guest users to the organization",
		Permission:           "User.Invite.All",
	}
	// DelegatedUserExportAll Export users' data
	DelegatedUserExportAll = Scope{
		AdminConsentRequired: true,
		Delegated:            true,
		Description:          "Allows the app to export an organizational user's data, when performed by a Company Administrator.",
		DisplayString:        "Export users' data",
		Permission:           "User.Export.All",
	}
)
