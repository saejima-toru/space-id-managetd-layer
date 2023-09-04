package spaces

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"identity-management/domains/id-store/models/spaces"
	"identity-management/domains/id-store/models/spaces/services/parameters"
	"identity-management/infrastructure/spaces/presentation"
)

// AwsCognitoSpacesRepository AWS Cognitoを使用したスペースリポジトリ具象
type AwsCognitoSpacesRepository struct {
	client             *cognitoidentityprovider.Client
	spacesPresentation presentation.SpacesRepositoryPresentation
}

func NewAwsCognitoSpacesRepository(
	spacesPresentation presentation.SpacesRepositoryPresentation) *AwsCognitoSpacesRepository {
	cgf, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		println("Init -> NewAwsCognitoSpacesRepository: ", err.Error())
	}

	client := cognitoidentityprovider.NewFromConfig(cgf)
	return &AwsCognitoSpacesRepository{
		client:             client,
		spacesPresentation: spacesPresentation,
	}
}

// CreateSpaces スペースの永続化保存 (Space == UserPool)
func (a *AwsCognitoSpacesRepository) CreateSpaces(createSpace spaces.Spaces) error {
	castSpace, isSpace := createSpace.(spaces.Space)
	if !isSpace {
		return errors.New("スペースの永続化に失敗しました。")
	}

	nameSpace := castSpace.NameSpace()
	createUserPoolInput := &cognitoidentityprovider.CreateUserPoolInput{
		PoolName: aws.String(nameSpace.String()),
	}

	_, createUserPoolErr := a.client.CreateUserPool(context.TODO(), createUserPoolInput)
	return createUserPoolErr
}

// DeleteSpace スペースの削除
func (a *AwsCognitoSpacesRepository) DeleteSpace(deleteSpace spaces.Spaces) error {
	castSpace, isSpace := deleteSpace.(spaces.Space)
	if !isSpace {
		return errors.New("スペースの削除に失敗しました。")
	}

	deleteSpaceId := castSpace.SpaceId()
	deleteUserPoolInput := &cognitoidentityprovider.DeleteUserPoolInput{
		UserPoolId: aws.String(deleteSpaceId.String()),
	}

	_, deleteErr := a.client.DeleteUserPool(context.TODO(), deleteUserPoolInput)
	return deleteErr
}

// FetchSpaces スペース情報を取得する
func (a *AwsCognitoSpacesRepository) FetchSpaces(
	parameter parameters.FetchSpacesParameter) spaces.Spaces {
	userPoolId := parameter.SpaceId()

	describeUserPoolInput := &cognitoidentityprovider.DescribeUserPoolInput{
		UserPoolId: aws.String(userPoolId.String()),
	}
	userPoolOutput, _ := a.client.DescribeUserPool(context.TODO(), describeUserPoolInput)
	return a.spacesPresentation.Presentation(userPoolOutput)
}
