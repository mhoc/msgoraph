package scopes

var (
	// DelegatedDeviceManagementAppsReadAll Read Microsoft Intune apps
	DelegatedDeviceManagementAppsReadAll = Scope{
		AdminConsentRequired: true,
		Description:          "Allows the app to read the properties, group assignments and status of apps, app configurations and app protection policies managed by Microsoft Intune.",
		DisplayString:        "Read Microsoft Intune apps",
		Permission:           "DeviceManagementApps.Read.All",
		Type:                 PermissionTypeDelegated,
	}
	// DelegatedDeviceManagementAppsReadWriteAll Read and write Microsoft Intune apps
	DelegatedDeviceManagementAppsReadWriteAll = Scope{
		AdminConsentRequired: true,
		Description:          "Allows the app to read and write the properties, group assignments and status of apps, app configurations and app protection policies managed by Microsoft Intune.",
		DisplayString:        "Read and write Microsoft Intune apps",
		Permission:           "DeviceManagementApps.ReadWrite.All",
		Type:                 PermissionTypeDelegated,
	}
	// DelegatedDeviceManagementConfigurationReadAll Read Microsoft Intune device configuration and policies
	DelegatedDeviceManagementConfigurationReadAll = Scope{
		AdminConsentRequired: true,
		Description:          "Allows the app to read properties of Microsoft Intune-managed device configuration and device compliance policies and their assignment to groups.",
		DisplayString:        "Read Microsoft Intune device configuration and policies",
		Permission:           "DeviceManagementConfiguration.Read.All",
		Type:                 PermissionTypeDelegated,
	}
	// DelegatedDeviceManagementConfigurationReadWriteAll "Read and write Microsoft Intune device configuration and policies"
	DelegatedDeviceManagementConfigurationReadWriteAll = Scope{
		AdminConsentRequired: true,
		Description:          "Allows the app to read and write properties of Microsoft Intune-managed device configuration and device compliance policies and their assignment to groups.",
		DisplayString:        "Read and write Microsoft Intune device configuration and policies",
		Permission:           "DeviceManagementConfiguration.ReadWrite.All",
		Type:                 PermissionTypeDelegated,
	}
	// DelegatedDeviceManagementManagedDevicesPrivilegedOperationsAll Perform user-impacting remote actions on Microsoft Intune devices
	DelegatedDeviceManagementManagedDevicesPrivilegedOperationsAll = Scope{
		AdminConsentRequired: true,
		Description:          "Allows the app to perform remote high impact actions such as wiping the device or resetting the passcode on devices managed by Microsoft Intune.",
		DisplayString:        "Perform user-impacting remote actions on Microsoft Intune devices",
		Permission:           "DeviceManagementManagedDevices.PrivilegedOperations.All",
		Type:                 PermissionTypeDelegated,
	}
	// DelegatedDeviceManagementManagedDevicesReadAll Read Microsoft Intune devices
	DelegatedDeviceManagementManagedDevicesReadAll = Scope{
		AdminConsentRequired: true,
		Description:          "Allows the app to read the properties of devices managed by Microsoft Intune.",
		DisplayString:        "Read Microsoft Intune devices",
		Permission:           "DeviceManagementManagedDevices.Read.All",
		Type:                 PermissionTypeDelegated,
	}
	// DelegatedDeviceManagementManagedDevicesReadWriteAll Read and write Microsoft Intune devices
	DelegatedDeviceManagementManagedDevicesReadWriteAll = Scope{
		AdminConsentRequired: true,
		Description:          "Allows the app to read and write the properties of devices managed by Microsoft Intune. Does not allow high impact operations such as remote wipe and password reset on the device’s owner.",
		DisplayString:        "Read and write Microsoft Intune devices",
		Permission:           "DeviceManagementManagedDevices.ReadWrite.All",
		Type:                 PermissionTypeDelegated,
	}
	// DelegatedDeviceManagementRBACReadAll Read Microsoft Intune RBAC settings
	DelegatedDeviceManagementRBACReadAll = Scope{
		AdminConsentRequired: true,
		Description:          "Allows the app to read the properties relating to the Microsoft Intune Role-Based Access Control (RBAC) settings.",
		DisplayString:        "Read Microsoft Intune RBAC settings",
		Permission:           "DeviceManagementRBAC.Read.All",
		Type:                 PermissionTypeDelegated,
	}
	// DelegatedDeviceManagementRBACReadWriteAll Read and write Microsoft Intune RBAC settings
	DelegatedDeviceManagementRBACReadWriteAll = Scope{
		AdminConsentRequired: true,
		Description:          "Allows the app to read and write the properties relating to the Microsoft Intune Role-Based Access Control (RBAC) settings.",
		DisplayString:        "Read and write Microsoft Intune RBAC settings",
		Permission:           "DeviceManagementRBAC.ReadWrite.All",
		Type:                 PermissionTypeDelegated,
	}
	// DelegatedDeviceManagementServiceConfigReadAll Read Microsoft Intune configuration
	DelegatedDeviceManagementServiceConfigReadAll = Scope{
		AdminConsentRequired: true,
		Description:          "Allows the app to read Intune service properties including device enrollment and third party service connection configuration.",
		DisplayString:        "Read Microsoft Intune configuration",
		Permission:           "DeviceManagementServiceConfig.Read.All",
		Type:                 PermissionTypeDelegated,
	}
	// DelegatedDeviceManagementServiceConfigReadWriteAll Read and write Microsoft Intune configuration
	DelegatedDeviceManagementServiceConfigReadWriteAll = Scope{
		AdminConsentRequired: true,
		Description:          "Allows the app to read and write Microsoft Intune service properties including device enrollment and third party service connection configuration.",
		DisplayString:        "Read and write Microsoft Intune configuration",
		Permission:           "DeviceManagementServiceConfig.ReadWrite.All",
		Type:                 PermissionTypeDelegated,
	}
)
