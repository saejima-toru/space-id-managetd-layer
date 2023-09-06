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

func NewUserSpacesManagedService(
	spacesRepository services.SpacesRepository) *UserSpacesManagedService {

	return &UserSpacesManagedService{
		spacesRepository: spacesRepository,
	}
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
	fetchSpaceBuilder := parameters.NewFetchSpaceParameterBuilder()
	fetchSpaceBuilder.SpaceId(spaceId.String())

	deleteSpace := u.spacesRepository.FetchSpaces(fetchSpaceBuilder.Build())
	if deleteSpace == nil {
		return errors.New("削除するスペースが存在しないため、スペースの削除を中断します。")
	}

	return u.spacesRepository.DeleteSpace(deleteSpace)
}

// ExistsSpace 既存にユーザースペースが存在しないかを確認する
func (u *UserSpacesManagedService) ExistsSpace(findSpaces spaces.Spaces) bool {
	var isSpace bool
	var findSpace spaces.Space
	fetchSpaceBuilder := parameters.NewFetchSpaceParameterBuilder()

	findSpace, isSpace = findSpaces.(spaces.Space)
	if isSpace == false {
		return false
	}
	findSpaceName := findSpace.NameSpace()
	fetchSpaceBuilder.NameSpace(findSpaceName.String())

	return u.spacesRepository.FetchSpaces(fetchSpaceBuilder.Build()) != nil
}
