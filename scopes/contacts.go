package scopes

var (
	// ApplicationContactsRead Read contacts in all mailboxes
	ApplicationContactsRead = Scope{
		AdminConsentRequired: true,
		Description:          "Allows the app to read all contacts in all mailboxes without a signed-in user.",
		DisplayString:        "Read contacts in all mailboxes",
		Permission:           "Contacts.Read",
		Type:                 PermissionTypeApplication,
	}
	// ApplicationContactsReadWrite Read and write contacts in all mailboxes
	ApplicationContactsReadWrite = Scope{
		AdminConsentRequired: true,
		Description:          "Allows the app to create, read, update, and delete all contacts in all mailboxes without a signed-in user.",
		DisplayString:        "Read and write contacts in all mailboxes",
		Permission:           "Contacts.ReadWrite",
		Type:                 PermissionTypeApplication,
	}
	// DelegatedContactsRead Read user contacts
	DelegatedContactsRead = Scope{
		Description:   "Allows the app to read user contacts.",
		DisplayString: "Read user contacts",
		Permission:    "Contacts.Read",
		Type:          PermissionTypeDelegated,
	}
	// DelegatedContactsReadShared Read user and shared contacts
	DelegatedContactsReadShared = Scope{
		Description:   "Allows the app to read contacts that the user has permissions to access, including the user's own and shared contacts.",
		DisplayString: "Read user and shared contacts",
		Permission:    "Contacts.Read.Shared",
		Type:          PermissionTypeDelegated,
	}
	// DelegatedContactsReadWrite Have full access to user contacts
	DelegatedContactsReadWrite = Scope{
		Description:   "Allows the app to create, read, update, and delete user contacts.",
		DisplayString: "Have full access to user contacts",
		Permission:    "Contacts.ReadWrite",
		Type:          PermissionTypeDelegated,
	}
	// DelegatedContactsReadWriteShared Read and write user and shared contacts
	DelegatedContactsReadWriteShared = Scope{
		Description:   "Allows the app to create, read, update and delete contacts that the user has permissions to, including the user's own and shared contacts.",
		DisplayString: "Read and write user and shared contacts",
		Permission:    "Contacts.ReadWrite.Shared",
		Type:          PermissionTypeDelegated,
	}
)
