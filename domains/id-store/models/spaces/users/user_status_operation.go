package users

import (
	"errors"
	"identity-management/domains/id-store/values/users/attributes/attribute"
)

// DisableAccess ユーザーを無効化する
func (u *User) DisableAccess() error {
	disabledStatus, err := u.userStatus.ShitToUserDisabledStatus()

	if disabledStatus != nil {
		u.userStatus = *disabledStatus
	}
	return err
}

// EnabledAccess ユーザのアクセスを有効化する
func (u *User) EnabledAccess() error {
	enabledUserStatus, err := u.userStatus.ShitToUserEnabledStatus()

	if enabledUserStatus != nil {
		u.userStatus = *enabledUserStatus
	}
	return err
}

// UpdateUserAttributes ユーザー属性を更新する
func (u *User) UpdateUserAttributes(
	updateAttributeRecords []attribute.UserAttributeRecord) error {
	if !u.userStatus.IsUserConfirmedStatus() {
		return errors.New("ユーザー属性の更新には、ユーザーの確認処理が終わっている必要があります。")
	}
	margeToAttributes := u.userAttributes
	afterMargeAttributes, err := margeToAttributes.MargeUserAttribute(updateAttributeRecords)

	if afterMargeAttributes != nil {
		u.userAttributes = *afterMargeAttributes
	}
	return err
}
