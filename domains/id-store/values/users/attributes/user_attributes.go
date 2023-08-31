package attributes

import (
	"errors"
	"identity-management/domains/id-store/values/users/attributes/attribute"
)

type UserAttributes struct {
	userAttributeRecords []attribute.UserAttributeRecord
}

func NewUserAttributes(userAttributeRecords []attribute.UserAttributeRecord) (*UserAttributes, error) {
	if len(userAttributeRecords) > 50 {
		return nil, errors.New("ユーザー属性は、最大50個まで追加することができます。")
	}
	if duplicate(userAttributeRecords) {
		return nil, errors.New("同じユーザー属性名を追加することはできません。")
	}

	return &UserAttributes{
		userAttributeRecords: userAttributeRecords,
	}, nil
}

// MargeUserAttribute ユーザー属性をマージする
func (u *UserAttributes) MargeUserAttribute(
	userAttrs []attribute.UserAttributeRecord) (*UserAttributes, error) {
	var err error
	to := u.userAttributeRecords
	for _, userAttr := range userAttrs {
		if idx := u.indexOf(userAttr); idx != -1 {
			err = to[idx].UpdateUserAttributeValue(*userAttr.AttributeValue())
			continue
		}
		return nil, errors.New("設定できないユーザー属性が指定されました。")
	}
	if err == nil {
		return NewUserAttributes(to)
	}
	return nil, err
}

// AddUserAttributeRecord ユーザー属性を追加する
func (u *UserAttributes) AddUserAttributeRecord(
	addAttrRecord attribute.UserAttributeRecord) (*UserAttributes, error) {

	addedAttributes := append(u.userAttributeRecords, addAttrRecord)
	return NewUserAttributes(addedAttributes)
}
