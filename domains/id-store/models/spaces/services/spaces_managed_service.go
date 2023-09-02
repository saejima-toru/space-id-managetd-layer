package services

import (
	"errors"
	"identity-management/domains/id-store/models/spaces"
	"identity-management/domains/id-store/models/spaces/services/parameters"
	spaceValue "identity-management/domains/id-store/values/spaces"
)

/**
ユーザープール管理サービス
*/

// SpacesManagesManagedService スペース管理サービス
type SpacesManagesManagedService interface {

	// ListSpaces ユーザープールリストを取得する
	ListSpaces() (spaces.Spaces, error)

	// CreateSpaces ユーザープールを作成する
	CreateSpaces(param parameters.SpacesParameter) error

	// DeleteSpace ユーザープールを削除する
	DeleteSpace(spaceId spaceValue.SpaceId) error

	// ExistsSpace 既存にスペースが存在しないかを確認する
	ExistsSpace(findSpace spaces.Spaces) bool
}

// UserSpacesManagesService ユーザースペース管理サービス
type UserSpacesManagesService struct {
	spacesRepository SpacesRepository
}

// ListSpaces ユーザースペースのリスト取得
func (u *UserSpacesManagesService) ListSpaces() (spaces.Spaces, error) {
	//TODO implement me
	panic("implement me")
}

// CreateSpaces ユーザースペースの作成
func (u *UserSpacesManagesService) CreateSpaces(param parameters.SpacesParameter) error {
	createSpaceFactory := spaces.NewCreateUserSpacesFactory()

	newUserSpace, factoryErr := createSpaceFactory.CreateNewSpace(param)
	if factoryErr != nil {
		return factoryErr
	}

	if u.ExistsSpace(newUserSpace) {
		return errors.New("既存に存在するスペース構成が指定されたため、作成することができません。")
	}
	return u.spacesRepository.CreateSpaces(newUserSpace)
}

// DeleteSpace ユーザースペースの削除
func (u *UserSpacesManagesService) DeleteSpace(spaceId spaceValue.SpaceId) error {
	//TODO implement me
	panic("implement me")
}

// ExistsSpace 既存にユーザースペースが存在しないかを確認する
func (u *UserSpacesManagesService) ExistsSpace(findSpace spaces.Spaces) bool {
	// TODO: 検索取得パラメーターを生成する
	fetchSpaceParams := &parameters.FetchSpacesParameter{}

	return u.spacesRepository.FetchSpaces(*fetchSpaceParams) != nil
}
