package users

// AutomaticRepliesSetting configuration settings to automatically notify the sender of an
// incoming email with a message from the signed-in user. For example, an automatic reply to
// notify that the signed-in user is unavailable to respond to emails.
type AutomaticRepliesSetting struct {
	ExternalAudience       string `json:"externalAudience"`
	ExternalReplyMessage   string `json:"externalReplyMessage"`
	InternalReplyMessage   string `json:"InternalReplyMessage"`
	ScheduledEndDateTime   string `json:"scheduledEndDateTime"`
	ScheduledStartDateTime string `json:"scheduledStartDateTime"`
	Status                 string `json:"status"`
}

// MailboxSettings Settings for the primary mailbox of the signed-in user.
type MailboxSettings struct {
	AutomaticRepliesSetting AutomaticRepliesSetting `json:"automaticRepliesSetting"`
	Language                LocaleInfo              `json:"localeInfo"`
	TimeZone                string                  `json:"timeZone"`
}
