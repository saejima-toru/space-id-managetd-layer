package services

import (
	"identity-management/domains/id-store/models/spaces"
	"identity-management/domains/id-store/models/spaces/services/parameters"
	spaceValue "identity-management/domains/id-store/values/spaces"
)

/**
ユーザープール管理サービス
*/

// UserPoolManagedService ユーザープール管理サービス
type UserPoolManagedService interface {

	// ListUserPools ユーザープールリストを取得する
	ListUserPools() (spaces.Spaces, error)

	// CreateUserPool ユーザープールを作成する
	CreateUserPool(param parameters.UserPoolParameter) error

	// DeleteUserPool ユーザープールを削除する
	DeleteUserPool(spaceId spaceValue.SpaceId) error
}
