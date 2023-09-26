package postmark

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

type EmailService struct {
	APIKey string
}

type EmailResponse struct {
	To          string `json:"to"`
	SubmittedAt string `json:"SubmittedAt"`
	MessageID   string `json:"MessageID"`
	ErrorCode   int    `json:"ErrorCode"`
	Message     string `json:"Message"`
}

func (e EmailService) SendActivationEmail(email string) error {

	body := []byte(fmt.Sprintf(`{
		"From": "hello@anthonygiang.xyz",
		"To": "hello@anthonygiang.xyz",
		"Subject": "Hello from Postmark",
		"HtmlBody": "%s",
		"MessageStream": "outbound",
	}`, email))

	r, err := http.NewRequest("POST", "https://api.postmarkapp.com/email", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal("Error creating request")
	}
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Postmark-Server-Token", e.APIKey)

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Fatal("Error sending request")
	}
	defer res.Body.Close()

	fmt.Printf("%+v", res)
	fmt.Printf("Sent the email")

	return nil
}
