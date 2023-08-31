package services

import (
	"identity-management/domains/id-store/models/users/services/parameters"
	"identity-management/domains/id-store/values/users"
)

// UserService ユーザーサービス
type UserService interface {

	// DisabledAccessUser ユーザーのアクセス無効化
	DisabledAccessUser(userId users.UserId) error

	// DeleteUser ユーザー削除
	DeleteUser(userId users.UserId) error

	// UpdateUser ユーザー情報のアップデート
	UpdateUser(userId users.UserId, userParameter parameters.UserParameter) error
}
