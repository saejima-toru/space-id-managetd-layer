package spaces

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"identity-management/domains/id-store/models/spaces"
	"identity-management/domains/id-store/models/spaces/services/parameters"
)

// AwsCognitoSpacesRepository AWS Cognitoを使用したスペースリポジトリ具象
type AwsCognitoSpacesRepository struct {
	client *cognitoidentityprovider.Client
}

func NewAwsCognitoSpacesRepository() *AwsCognitoSpacesRepository {
	cgf, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		println("Init -> NewAwsCognitoSpacesRepository: ", err.Error())
	}

	client := cognitoidentityprovider.NewFromConfig(cgf)
	return &AwsCognitoSpacesRepository{
		client: client,
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

	if _, createUserPoolErr := a.client.CreateUserPool(context.TODO(), createUserPoolInput); createUserPoolErr != nil {
		return createUserPoolErr
	}
	return nil
}

func (a *AwsCognitoSpacesRepository) DeleteSpace(deleteSpace spaces.Spaces) error {
	//TODO implement me
	panic("implement me")
}

func (a *AwsCognitoSpacesRepository) FetchSpaces(parameter parameters.FetchSpacesParameter) spaces.Spaces {
	//TODO implement me
	panic("implement me")
}
