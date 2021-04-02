package twiliosms

// Factory is create new object
func Factory(AccountSID string, AuthToken string, MessageSID string) *TWILIOSMS {
	obj := new(TWILIOSMS)
	obj.AccountSID = AccountSID
	obj.AuthToken = AuthToken
	obj.MessageSID = MessageSID
	return obj
}
