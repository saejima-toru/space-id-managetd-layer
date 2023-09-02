package services

import (
	"identity-management/domains/id-store/models/spaces/services/parameters"
	"identity-management/domains/id-store/values/users"
)

// SpacesUserManagedService スペースユーザー管理サービス
type SpacesUserManagedService interface {

	// DisabledAccessUser ユーザーのアクセス無効化
	DisabledAccessUser(userId users.UserId) error

	// EnabledAccessUser ユーザーのアクセス有効化
	EnabledAccessUser(userId users.UserId) error

	// UpdateUser ユーザー情報のアップデート
	UpdateUser(userId users.UserId, userParameter parameters.UserParameter) error
}
