package twiliosms

// New is same Factory function.
func New1(AccountSID string, AuthToken string, MessageSID string) *TWILIOSMS {
	return Factory(AccountSID, AuthToken, MessageSID)
}
