package users

import (
	"errors"
	"identity-management/domains/id-store/values/users/attributes/attribute"
)

type UserAttributes struct {
	userAttributes []UserAttribute
}

func NewUserAttributes(userAttributes []UserAttribute) (*UserAttributes, error) {
	userAttrs := &UserAttributes{}
	for _, userAttr := range userAttributes {
		if userAttrs.contains(userAttr) {
			return nil, errors.New("同一のユーザー属性を含めることはできません。")
		}
		userAttrs.userAttributes = append(userAttrs.userAttributes, userAttr)
	}

	return userAttrs, nil
}

func (a *UserAttributes) AddUserAttribute(userAttribute UserAttribute) error {
	origin := a.userAttributes
	if a.contains(userAttribute) {
		return errors.New("既に追加されている属性のため追加できません。")
	}
	origin = append(origin, userAttribute)
	a.userAttributes = origin
	return nil
}

func (a *UserAttributes) UserAttributes() []UserAttribute {
	return a.userAttributes
}

func (a *UserAttributes) contains(userAttr UserAttribute) bool {
	for _, attr := range a.userAttributes {
		if attr.EqualTo(userAttr) {
			return true
		}
	}
	return false
}

type UserAttribute struct {
	userAttributeName  attribute.UserAttributeName  // アトリビュート名
	userAttributeValue attribute.UserAttributeValue // アトリビュート値
}

func NewUserAttribute(
	userAttributeName attribute.UserAttributeName,
	userAttributeValue attribute.UserAttributeValue) *UserAttribute {

	return &UserAttribute{
		userAttributeName:  userAttributeName,
		userAttributeValue: userAttributeValue,
	}
}

// EqualTo 同一のユーザー属性であるかを判定する
func (a *UserAttribute) EqualTo(userAttribute UserAttribute) bool {

	return a.userAttributeName.EqualTo(userAttribute.userAttributeName) &&
		a.userAttributeValue.EqualTo(userAttribute.userAttributeValue)
}
