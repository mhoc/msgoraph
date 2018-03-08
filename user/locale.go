package user

// LocaleInfo Information about the locale, including the preferred language and country/region, of the signed-in user.
type LocaleInfo struct {
	DisplayName string `json:"displayName"`
	Locale      string `json:"locale"`
}
