package users

import (
	"errors"
	"regexp"
)

type UserName struct {
	userName string
}

func NewUserName(userName string) (*UserName, error) {
	if len(userName) < 1 || len(userName) > 128 {
		return nil, errors.New("ユーザー名は、1〜128文字以内である必要があります。")
	}
	if r := regexp.MustCompile(`^.*[^a-zA-Z\d+=,.@-].*$`); r.MatchString(userName) {
		return nil, errors.New("ユーザー名は、大文字英字・番号(0〜9)・特殊文字(+ = , . @ -)のみ含めることができます。")
	}
	return &UserName{userName: userName}, nil
}
