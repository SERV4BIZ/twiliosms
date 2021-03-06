package twiliosms

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/SERV4BIZ/gfp/jsons"
)

// Send sms to number phone
func (me *TWILIOSMS) Send(txtToNumber string, txtMessage string) (string, error) {
	apiurl := "https://api.twilio.com/2010-04-01/Accounts/" + me.AccountSID + "/Messages.json"
	method := "POST"

	params := url.Values{}
	params.Add("To", txtToNumber)
	params.Add("MessagingServiceSid", me.MessageSID)
	params.Add("Body", txtMessage)
	payload := strings.NewReader(params.Encode())

	client := &http.Client{}
	req, err := http.NewRequest(method, apiurl, payload)

	if err != nil {
		return "", err
	}

	txtAuthen := base64.StdEncoding.EncodeToString([]byte(me.AccountSID + ":" + me.AuthToken))
	req.Header.Add("Authorization", "Basic "+txtAuthen)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	jsoReq, errReq := jsons.ObjectString(string(body))
	if errReq != nil {
		return "", errReq
	}

	if !jsoReq.CheckKey("sid") {
		return "", errors.New(jsoReq.ToString())
	}

	return jsoReq.GetString("sid"), nil
}
