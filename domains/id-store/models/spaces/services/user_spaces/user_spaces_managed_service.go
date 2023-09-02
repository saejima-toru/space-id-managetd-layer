package user_spaces

import (
	"errors"
	"identity-management/domains/id-store/models/spaces"
	"identity-management/domains/id-store/models/spaces/services"
	"identity-management/domains/id-store/models/spaces/services/parameters"
	spaceValue "identity-management/domains/id-store/values/spaces"
)

type UserSpacesManagedService struct {
	spacesRepository services.SpacesRepository
}

// ListSpaces ユーザースペースのリスト取得
func (u *UserSpacesManagedService) ListSpaces() (spaces.Spaces, error) {
	//TODO implement me
	panic("implement me")
}

// CreateSpaces ユーザースペースの作成
func (u *UserSpacesManagedService) CreateSpaces(param parameters.SpacesParameter) error {
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
func (u *UserSpacesManagedService) DeleteSpace(spaceId spaceValue.SpaceId) error {
	//TODO implement me
	panic("implement me")
}

// ExistsSpace 既存にユーザースペースが存在しないかを確認する
func (u *UserSpacesManagedService) ExistsSpace(findSpace spaces.Spaces) bool {
	// TODO: 検索取得パラメーターを生成する
	fetchSpaceParams := &parameters.FetchSpacesParameter{}

	return u.spacesRepository.FetchSpaces(*fetchSpaceParams) != nil
}
