package presentation

import (
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"identity-management/domains/id-store/models/spaces"
	spacesValue "identity-management/domains/id-store/values/spaces"
	"identity-management/domains/id-store/values/timestamp"
)

// UserSpacesRepositoryPresentation ユーザースペースリポジトリプレゼンター
type UserSpacesRepositoryPresentation struct{}

// Presentation ユーザースペースへ変換する
func (u *UserSpacesRepositoryPresentation) Presentation(
	output *cognitoidentityprovider.DescribeUserPoolOutput) spaces.Spaces {

	if output == nil {
		return nil
	}
	fetchUserPool := output.UserPool

	spaceId, spaceIdErr := spacesValue.NewSpaceId(*fetchUserPool.Id)
	if spaceIdErr != nil {
		return nil
	}
	nameSpace, nameSpaceId := spacesValue.NewNameSpace(*fetchUserPool.Name)
	if nameSpaceId != nil {
		return nil
	}
	createAtTime := fetchUserPool.CreationDate
	updateAtTime := fetchUserPool.LastModifiedDate

	spaceCreateAt := timestamp.NewTimeStampFromUnixTime(createAtTime.Unix())
	spaceUpdateAt := timestamp.NewTimeStampFromUnixTime(updateAtTime.Unix())
	return spaces.NewSpace(
		*spaceId,
		*nameSpace,
		*spaceCreateAt,
		*spaceUpdateAt,
	)
}
