package services

import (
	"identity-management/domains/id-store/models/spaces/users"
	"identity-management/domains/id-store/models/spaces/users/parameters"
	"identity-management/domains/id-store/values/spaces"
	valueUser "identity-management/domains/id-store/values/users"
)

/*
スペース内ユーザー管理サービス
*/

// SpacesUserManagedService スペースユーザー管理サービス
type SpacesUserManagedService interface {

	// CreateAuthenticateUser 認証ユーザーの作成
	CreateAuthenticateUser(params parameters.UserParams) error

	// DeleteUser ユーザーの削除
	DeleteUser(deleteUserId valueUser.UserId, spacesId spaces.SpaceId) error

	// DisabledAccessUser ユーザーのアクセス無効化
	DisabledAccessUser(userId valueUser.UserId, id spaces.SpaceId) error

	// EnabledAccessUser ユーザーのアクセス有効化
	EnabledAccessUser(userId valueUser.UserId, id spaces.SpaceId) error

	// ExistsUser ユーザーの存在を確認する
	ExistsUser(existsCheckUser users.User, id spaces.SpaceId) bool
}
