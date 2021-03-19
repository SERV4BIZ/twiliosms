package twiliosms

func New(AccountSID string, AuthToken string, MessageSID string) *TWILIOSMS {
	return Factory(AccountSID, AuthToken, MessageSID)
}
