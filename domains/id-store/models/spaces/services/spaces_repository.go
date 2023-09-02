package services

import (
	"identity-management/domains/id-store/models/spaces"
	"identity-management/domains/id-store/models/spaces/services/parameters"
)

// TODO: 実装

// SpacesRepository スペースリポジトリ
type SpacesRepository interface {

	// CreateSpaces スペース情報の永続保存
	CreateSpaces(createSpace spaces.Spaces) error

	// DeleteSpace スペースの削除
	DeleteSpace(deleteSpace spaces.Spaces) error

	// FetchSpaces スペースの取得
	FetchSpaces(parameter parameters.FetchSpacesParameter) spaces.Spaces
}
