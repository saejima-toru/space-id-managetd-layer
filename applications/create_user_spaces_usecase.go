package applications

import (
	"errors"
	"identity-management/domains/id-store/models/spaces/services"
	"identity-management/domains/id-store/models/spaces/services/parameters"
)

// CreateUserSpacesUseCase ユーザースペース作成ユースケース
type CreateUserSpacesUseCase struct {
	spacesService services.SpacesManagedService
}

// CreateSpace スペースの作成を実行する
func (c *CreateUserSpacesUseCase) CreateSpace(
	input *CreateSpacesInput) (*CreateSpacesOutput, error) {
	spacesParamBuilder := parameters.NewSpaceParameterBuilder()
	findSpaceParamBuilder := parameters.NewFetchSpaceParameterBuilder()

	spacesParamBuilder.NameSpace(input.CreateNameSpace)
	createSpaceErr := c.spacesService.CreateSpaces(spacesParamBuilder.Build())
	if createSpaceErr != nil {
		return nil, createSpaceErr
	}

	findSpaceParamBuilder.NameSpace(*input.CreateNameSpace)
	createdSpace := c.spacesService.ExistsSpace(findSpaceParamBuilder.Build())
	if createdSpace == false {
		return nil, errors.New("スペースの作成に失敗しました。")
	}

	return &CreateSpacesOutput{}, nil
}
