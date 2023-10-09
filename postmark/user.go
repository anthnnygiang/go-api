package postmark

import (
	"anthnnygiang/api-template/app"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type EmailService struct {
	HTTPClient *http.Client
	APIKey     string
}

func (es EmailService) SendActivationEmail(email app.ActivationEmail) (*app.ActivationEmail, error) {

	//Use values from email argument when ready
	body := []byte(fmt.Sprintf(`{
		"From": "hello@anthonygiang.xyz",
		"To": "hello@anthonygiang.xyz",
		"Subject": "Hello from Postmark",
		"HtmlBody": "Hi, your email is: %s",
		"MessageStream": "outbound",
	}`, email.To))

	//Be aware of token referer leakage
	r, err := http.NewRequest("POST", "https://api.postmarkapp.com/email", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Postmark-Server-Token", es.APIKey)

	res, err := es.HTTPClient.Do(r)
	if err != nil {
		return nil, err
	}
	resBody := &Response{}
	err = json.NewDecoder(res.Body).Decode(resBody)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	fmt.Printf("%+v", resBody)

	return &email, nil
}
