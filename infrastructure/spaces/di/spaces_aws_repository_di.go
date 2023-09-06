package di

import (
	"identity-management/infrastructure/spaces"
	"identity-management/infrastructure/spaces/presentation"
)

// SpacesAwsRepositoryDI AWSを使用する想定のスペースリポジトリDI
type SpacesAwsRepositoryDI struct{}

func NewUserSpacesRepositoryForAws() *spaces.AwsCognitoSpacesRepository {

	return spaces.NewAwsCognitoSpacesRepository(
		presentation.NewUserSpacesRepositoryPresentation(),
	)
}
