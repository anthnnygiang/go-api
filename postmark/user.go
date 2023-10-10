package postmark

import (
	"anthnnygiang/api-template/app"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type EmailService struct {
	HTTPClient *http.Client
	APIKey     string
}

func (es EmailService) SendActivationEmail(email app.ActivationEmail) (*app.ActivationEmail, error) {

	//Use values from email argument when ready
	payload := []byte(fmt.Sprintf(`{
		"From": "hello@anthonygiang.xyz",
		"To": "hello@anthonygiang.xyz",
		"Subject": "Activate your account",
		"HtmlBody": "true destination: %s, raw token: %s, copy paste into form.",
		"MessageStream": "outbound",
	}`, email.To, strings.ToUpper(email.RawToken)))

	//Be aware of token referer leakage
	req, err := http.NewRequest("POST", "https://api.postmarkapp.com/email", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Postmark-Server-Token", es.APIKey)

	res, err := es.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	var msg Response
	err = json.NewDecoder(res.Body).Decode(&msg)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	fmt.Printf("%+v\n", msg)

	return &email, nil
}
