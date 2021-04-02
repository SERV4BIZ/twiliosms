package twiliosms

// Create is same Factory function
func Create(AccountSID string, AuthToken string, MessageSID string) *TWILIOSMS {
	return Factory(AccountSID, AuthToken, MessageSID)
}
