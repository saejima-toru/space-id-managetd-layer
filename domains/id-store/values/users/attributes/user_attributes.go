package attributes

import (
	"errors"
)

type UserAttributes struct {
	userAttributeRecords []UserAttributeRecord
}

func NewUserAttributes(userAttributeRecords []UserAttributeRecord) (*UserAttributes, error) {
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

// AddUserAttributeRecord ユーザー属性を追加する
func (u *UserAttributes) AddUserAttributeRecord(
	addAttrRecord UserAttributeRecord) (*UserAttributes, error) {

	addedAttributes := append(u.userAttributeRecords, addAttrRecord)
	return NewUserAttributes(addedAttributes)
}

// UpdateUserAttributeValue 指定したユーザー属性値を変更する
func (u *UserAttributes) UpdateUserAttributeValue(
	attrName UserAttributeName, attrValue UserAttributeValue) error {
	targetAttrRecord := u.findByAttrName(attrName)

	if targetAttrRecord != nil {
		return targetAttrRecord.UpdateUserAttributeValue(attrValue)
	}
	return nil
}
