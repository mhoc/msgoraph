package scopes

var (
	// ApplicationDeviceReadWriteAll Read and write devices
	ApplicationDeviceReadWriteAll = Scope{
		AdminConsentRequired: true,
		Application:          true,
		Description:          "Allows the app to read and write all device properties without a signed in user. Does not allow device creation, device deletion, or update of device alternative security identifiers.",
		DisplayString:        "Read and write devices",
		Permission:           "Device.ReadWrite.All",
	}
	// DelegatedDeviceRead Read user devices
	DelegatedDeviceRead = Scope{
		Delegated:     true,
		Description:   "Allows the app to read a user's list of devices on behalf of the signed-in user.",
		DisplayString: "Read user devices",
		Permission:    "Device.Read",
	}
	// DelegatedDeviceCommand Communicate with user devices
	DelegatedDeviceCommand = Scope{
		Delegated:     true,
		Description:   "Allows the app to launch another app or communicate with another app on a user's device on behalf of the signed-in user.",
		DisplayString: "Communicate with user devices",
		Permission:    "Device.Command",
	}
)
