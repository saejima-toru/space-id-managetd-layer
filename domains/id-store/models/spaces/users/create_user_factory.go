package users

import (
	"errors"
	"identity-management/domains/id-store/models/spaces/users/parameters"
	"identity-management/domains/id-store/values/timestamp"
	"identity-management/domains/id-store/values/users"
	"identity-management/domains/id-store/values/users/attributes/attribute"
	"identity-management/domains/id-store/values/users/status"
)

type CreateUserFactory struct{}

func NewCreateUserFactory() *CreateUserFactory {

	return &CreateUserFactory{}
}

// CreateUser 新規ユーザーの作成を実行する
func (c *CreateUserFactory) CreateUser(params parameters.UserParams) (*User, error) {
	now := timestamp.NewTimeStamp()
	userName := params.UserName

	if params.UserId == nil {
		return nil, errors.New("ユーザーIDの指定は、必須です。")
	}
	if userName == nil {
		userName, _ = users.NewUserName(params.UserId.String())
	}
	userAttributes, err := c.createUserAttributes(params.UserAttributesParams)
	if err != nil {
		return nil, err
	}

	return NewUser(
		*params.UserId,
		*userName,
		*status.NewRegisteredUserConfirmStatus(),
		*userAttributes,
		now,
		now,
	), nil
}

// createUserAttributes　ユーザー属性を組み立てる
func (c *CreateUserFactory) createUserAttributes(
	params *parameters.UserAttributesParam) (*UserAttributes, error) {
	var userAttributes []UserAttribute

	for userAttributeName, userAttributeParam := range *params {
		newUserAttributeName, err := attribute.NewUserAttributeName(userAttributeName)
		if err != nil {
			return nil, err
		}

		userAttributes = append(userAttributes, *NewUserAttribute(
			*newUserAttributeName,
			*c.createUserAttributeValue(userAttributeParam),
		))
	}
	return NewUserAttributes(userAttributes)
}

// createUserAttributeValue ユーザー属性値を生成する
func (c *CreateUserFactory) createUserAttributeValue(
	valueParam parameters.UserAttributeValueParam) *attribute.UserAttributeValue {
	var value = ""

	if valueParam.S != nil {
		value = valueParam.S.String()
	}
	if valueParam.N != nil {
		value = valueParam.N.String()
	}
	return attribute.NewUserAttributeValue(value)
}
