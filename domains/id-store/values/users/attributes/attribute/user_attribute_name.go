package attribute

import "errors"

type UserAttributeName struct {
	userAttrName string
}

func NewUserAttributeName(userAttrName string) (*UserAttributeName, error) {
	if len(userAttrName) < 1 || len(userAttrName) > 20 {
		return nil, errors.New("ユーザー属性名は、1文字以上20以下である必要があります。")
	}

	return &UserAttributeName{userAttrName: userAttrName}, nil
}

func (u *UserAttributeName) EqualTo(v UserAttributeName) bool {
	return u.userAttrName == v.userAttrName
}

func (u *UserAttributeName) String() string {
	return u.userAttrName
}
