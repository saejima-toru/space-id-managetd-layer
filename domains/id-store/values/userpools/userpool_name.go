package userpools

import (
	"errors"
	"regexp"
)

type UserPoolName struct {
	userPoolName string
}

func NewUserPoolName(userPoolName string) (*UserPoolName, error) {
	if len(userPoolName) < 1 || len(userPoolName) > 128 {
		return nil, errors.New("ユーザープール名は、1〜128文字である必要があります。")
	}
	if validate := regexp.MustCompile("^.*[^a-zA-Z\\d+=,.@-].*$"); validate.MatchString(userPoolName) {
		return nil, errors.New("ユーザー名は、大文字英字・番号(0〜9)・特殊文字(+ = , . @ -)のみ含めることができます。")
	}

	return &UserPoolName{
		userPoolName: userPoolName,
	}, nil
}

func (p *UserPoolName) String() string {
	return p.userPoolName
}
