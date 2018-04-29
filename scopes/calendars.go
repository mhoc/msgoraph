package scopes

var (
	// ApplicationCalendarsRead Read calendars in all mailboxes
	ApplicationCalendarsRead = Scope{
		AdminConsentRequired: true,
		Application:          true,
		Description:          "Allows the app to read events of all calendars without a signed-in user.",
		DisplayString:        "Read calendars in all mailboxes",
		Permission:           "Calendars.Read",
	}
	// ApplicationCalendarsReadWrite Read and write calendars in all mailboxes
	ApplicationCalendarsReadWrite = Scope{
		AdminConsentRequired: true,
		Application:          true,
		Description:          "Allows the app to create, read, update, and delete events of all calendars without a signed-in user.",
		DisplayString:        "Read and write calendars in all mailboxes",
		Permission:           "Calendars.ReadWrite",
	}
	// DelegatedCalendarsRead Read user calendars
	DelegatedCalendarsRead = Scope{
		Delegated:     true,
		Description:   "Allows the app to read events in user calendars.",
		DisplayString: "Read user calendars",
		Permission:    "Calendars.Read",
	}
	// DelegatedCalendarsReadShared Read user and shared calendars
	DelegatedCalendarsReadShared = Scope{
		Delegated:     true,
		Description:   "Allows the app to read events in all calendars that the user can access, including delegate and shared calendars.",
		DisplayString: "Read user and shared calendars",
		Permission:    "Calendars.Read.Shared",
	}
	// DelegatedCalendarsReadWrite Have full access to user calendars
	DelegatedCalendarsReadWrite = Scope{
		Delegated:     true,
		Description:   "Allows the app to create, read, update, and delete events in user calendars.",
		DisplayString: "Have full access to user calendars",
		Permission:    "Calendars.ReadWrite",
	}
	// DelegatedCalendarsReadWriteShared Read and write user and shared calendars
	DelegatedCalendarsReadWriteShared = Scope{
		Delegated:     true,
		Description:   "Allows the app to create, read, update and delete events in all calendars the user has permissions to access. This includes delegate and shared calendars.",
		DisplayString: "Read and write user and shared calendars",
		Permission:    "Calendars.ReadWrite.Shared",
	}
)
