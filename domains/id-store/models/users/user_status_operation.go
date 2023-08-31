package users

import (
	"errors"
	"identity-management/domains/id-store/values/users/attributes"
)

// DisableAccess ユーザーを無効化する
func (u *User) DisableAccess() error {
	disabledStatus, err := u.userStatus.ShitToUserDisabledStatus()

	if disabledStatus != nil {
		u.userStatus = *disabledStatus
	}
	return err
}

// UpdateUserAttributes ユーザー属性を更新する
func (u *User) UpdateUserAttributes(
	margeUserAttributes attributes.UserAttributes) error {
	if !u.userStatus.IsUserConfirmedStatus() {
		return errors.New("ユーザー確認が終わっていないため、ユーザー属性の更新に失敗しました。")
	}
	return nil
}
