package spaces

import (
	"identity-management/domains/id-store/models/spaces"
	"identity-management/domains/id-store/models/spaces/services/parameters"
)

// AwsCognitoSpacesRepository AWS Cognitoを使用したスペースリポジトリ具象
type AwsCognitoSpacesRepository struct{}

func (a *AwsCognitoSpacesRepository) CreateSpaces(createSpace spaces.Spaces) error {
	//TODO implement me
	panic("implement me")
}

func (a *AwsCognitoSpacesRepository) DeleteSpace(deleteSpace spaces.Spaces) error {
	//TODO implement me
	panic("implement me")
}

func (a *AwsCognitoSpacesRepository) FetchSpaces(parameter parameters.FetchSpacesParameter) spaces.Spaces {
	//TODO implement me
	panic("implement me")
}
