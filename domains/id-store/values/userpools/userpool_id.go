package userpools

import "errors"

type UserPoolId struct {
	userPoolId string
}

func NewUserPoolId(userPoolId string) (*UserPoolId, error) {
	if len(userPoolId) < 1 {
		return nil, errors.New("ユーザープールIDの指定は、必須です。")
	}

	return &UserPoolId{
		userPoolId: userPoolId,
	}, nil
}

// EqualTo 同一ユーザープールIDであるかを比較する
func (u *UserPoolId) EqualTo(compareUserPoolId UserPoolId) bool {
	return u.String() == compareUserPoolId.String()
}

func (u *UserPoolId) String() string {
	return u.userPoolId
}
