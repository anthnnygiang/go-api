package postmark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type EmailService struct {
	HTTPClient *http.Client
	APIKey     string
}

type Response struct {
	To          string    `json:"To"`
	SubmittedAt time.Time `json:"SubmittedAt"`
	MessageID   string    `json:"MessageID"`
	ErrorCode   int       `json:"ErrorCode"`
	Message     string    `json:"Message"`
}

func (e EmailService) SendActivationEmail(email string) (*Response, error) {

	body := []byte(fmt.Sprintf(`{
		"From": "hello@anthonygiang.xyz",
		"To": "hello@anthonygiang.xyz",
		"Subject": "Hello from Postmark",
		"HtmlBody": "Hi, your email is: %s",
		"MessageStream": "outbound",
	}`, email))

	r, err := http.NewRequest("POST", "https://api.postmarkapp.com/email", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal("Error creating request")
	}
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Postmark-Server-Token", e.APIKey)

	res, err := e.HTTPClient.Do(r)
	if err != nil {
		log.Fatal("Error sending request")
	}
	resBody := &Response{}
	err = json.NewDecoder(res.Body).Decode(resBody)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	fmt.Printf("%+v", resBody)

	return resBody, nil
}
