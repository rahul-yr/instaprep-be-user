package login

import (
	"errors"
	"net/mail"
	"strings"
)

var (
	ALLOWED_EMAIL_PROVIDERS = map[string]bool{
		"gmail.com":   true,
		"outlook.com": true,
		"apple.com":   true}
)

type Response struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type RequestParams struct {
	Email string `json:"email"  xml:"email" form:"email"`
}

type PubsubEmailMessage struct {
	Email   string                 `json:"email"`
	Type    string                 `json:"type"`
	Payload map[string]interface{} `json:"payload"`
}

func (rp *RequestParams) Validate() error {
	email_id, err := normalizeEmail(rp.Email)
	if err != nil {
		return err
	}
	rp.Email = email_id
	return nil
}

func normalizeEmail(email string) (string, error) {
	email = strings.TrimSpace(email)
	email = strings.ToLower(email)
	_, err := mail.ParseAddress(email)
	if err != nil {
		return "", err
	}
	s := strings.Split(email, "@")
	if _, ok := ALLOWED_EMAIL_PROVIDERS[s[len(s)-1]]; ok {
		return email, nil
	}
	return "", errors.New("email : Your Email provider is not yet supported")
}
