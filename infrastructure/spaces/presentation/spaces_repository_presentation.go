package presentation

import (
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"identity-management/domains/id-store/models/spaces"
)

// SpacesRepositoryPresentation スペース：リポジトリプレゼンター
type SpacesRepositoryPresentation interface {

	// Presentation ドメインモデルに変換する
	Presentation(output *cognitoidentityprovider.DescribeUserPoolOutput) spaces.Spaces
}
