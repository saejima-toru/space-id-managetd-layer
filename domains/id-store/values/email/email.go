package email

import (
	"errors"
	"regexp"
)

type Email struct {
	email string
}

func NewEmail(email string) (*Email, error) {
	if len(email) < 1 || len(email) > 319 {
		return nil, errors.New("メールアドレスは、1〜319文字以内である必要があります。")
	}
	if r := regexp.MustCompile(`^[a-zA-Z\d_.-]+@[a-zA-Z\d-]+\.[a-zA-Z]{2,6}$`); !r.MatchString(email) {
		return nil, errors.New("メールアドレスの形式に問題があります。")
	}
	return &Email{email: email}, nil
}

func (e *Email) String() string {
	return e.email
}
