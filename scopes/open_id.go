package scopes

var (
	// DelegatedEmail View users' email address
	DelegatedEmail = Scope{
		Delegated:     true,
		Description:   "Allows the app to read your users' primary email address.",
		DisplayString: "View users' email address",
		Permission:    "email",
	}
	// DelegatedOfflineAccess Access user's data anytime
	DelegatedOfflineAccess = Scope{
		Delegated:     true,
		Description:   "Allows the app to read and update user data, even when they are not currently using the app.",
		DisplayString: "Access user's data anytime",
		Permission:    "offline_access",
	}
	// DelegatedOpenID Sign users in
	DelegatedOpenID = Scope{
		Delegated:     true,
		Description:   "Allows users to sign in to the app with their work or school accounts and allows the app to see basic user profile information.",
		DisplayString: "Sign users in",
		Permission:    "openid",
	}
	// DelegatedProfile View users' basic profile
	DelegatedProfile = Scope{
		Delegated:     true,
		Description:   "Allows the app to see your users' basic profile (name, picture, user name).",
		DisplayString: "View users' basic profile",
		Permission:    "profile",
	}
)
