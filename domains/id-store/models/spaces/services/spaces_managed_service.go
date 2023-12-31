package services

import (
	"identity-management/domains/id-store/models/spaces"
	"identity-management/domains/id-store/models/spaces/services/parameters"
	spaceValue "identity-management/domains/id-store/values/spaces"
)

/**
ユーザープール管理サービス
*/

// SpacesManagedService スペース管理サービス
type SpacesManagedService interface {

	// CreateSpaces ユーザープールを作成する
	CreateSpaces(param parameters.SpacesParameter) error

	// DeleteSpace ユーザープールを削除する
	DeleteSpace(spaceId spaceValue.SpaceId) error

	// ExistsSpace 既存にスペースが存在しないかを確認する
	ExistsSpace(findSpace spaces.Spaces) bool
}
