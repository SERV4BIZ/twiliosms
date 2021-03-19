package twiliosms

func Factory(AccountSID string, AuthToken string, MessageSID string) *TWILIOSMS {
	obj := new(TWILIOSMS)
	obj.AccountSID = AccountSID
	obj.AuthToken = AuthToken
	obj.MessageSID = MessageSID
	return obj
}
