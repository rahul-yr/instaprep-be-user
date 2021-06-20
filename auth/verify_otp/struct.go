package verifyotp

import (
	"errors"
	"net/mail"
	"strconv"
	"strings"
)

var (
	ALLOWED_EMAIL_PROVIDERS = map[string]bool{
		"gmail.com":   true,
		"outlook.com": true,
		"apple.com":   true}
)

type Response struct {
	Token   string `json:"token"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type RequestParams struct {
	Email string `json:"email"  xml:"email" form:"email"`
	OTP   int    `json:"otp"  xml:"otp" form:"otp"`
}

func (rp *RequestParams) Validate() error {
	// email validation
	email_id, err := normalizeEmail(rp.Email)
	if err != nil {
		return err
	}
	rp.Email = email_id
	// otp validation
	err = checkOTP(rp.OTP)
	if err != nil {
		return err
	}
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

func checkOTP(otp int) error {
	if len(strconv.Itoa(otp)) != 6 {
		return errors.New("otp : OTP is not valid")
	}
	return nil
}
