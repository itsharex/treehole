package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"
)

const reUrl = "https://www.recaptcha.net/recaptcha/api/siteverify"

type recaptchaResponse struct {
	Score       float64   `json:"score"`
	Success     bool      `json:"success"`
	ChallengeTs time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	Action      string    `json:"action"`
	ErrorCodes  []string  `json:"error-codes"`
}

func Recaptcha(token string) error {
	formValue := url.Values{}
	formValue.Set("secret", recaptchaSec)
	formValue.Set("response", token)
	formStr := formValue.Encode()
	formBytes := []byte(formStr)
	response, err := http.Post(reUrl, "application/x-www-form-urlencoded", bytes.NewReader(formBytes))
	if err != nil {
		return err
	}
	r, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	var resp recaptchaResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return err
	}
	if !resp.Success {
		return errors.New("recaptcha failed")
	}
	return nil
}
