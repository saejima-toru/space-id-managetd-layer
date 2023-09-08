package users

import "errors"

type Password struct {
	password string
}

func NewPassword(password string) (*Password, error) {
	if len(password) < 1 {
		return nil, errors.New("パスワードの指定は、必須です。")
	}

	return &Password{
		password: password,
	}, nil
}

func (p *Password) String() string {
	return p.password
}
