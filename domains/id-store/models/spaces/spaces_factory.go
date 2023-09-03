package spaces

import (
	"github.com/google/uuid"
	"identity-management/domains/id-store/models/spaces/services/parameters"
	"identity-management/domains/id-store/values/spaces"
	datetime "identity-management/domains/id-store/values/timestamp"
)

type CreateSpacesFactory interface {

	// CreateNewSpace 新規スペースを構成する
	CreateNewSpace(parameter parameters.SpacesParameter) (Spaces, error)
}

// CreateUserSpacesFactory ユーザースペースファクトリ
type CreateUserSpacesFactory struct{}

func NewCreateUserSpacesFactory() *CreateUserSpacesFactory {
	return &CreateUserSpacesFactory{}
}

func (c *CreateUserSpacesFactory) CreateNewSpace(
	parameter parameters.SpacesParameter) (Spaces, error) {
	timestamp := datetime.NewTimeStamp()

	if parameter.NameSpace() != nil {
		return nil, nil
	}
	return NewSpace(
		*generateSpaceId(),
		*parameter.NameSpace(),
		timestamp,
		timestamp,
	), nil
}

func generateSpaceId() *spaces.SpaceId {
	spaceUuId, err := uuid.NewUUID()
	if err != nil {
		return generateSpaceId()
	}
	spaceId, _ := spaces.NewSpaceId(spaceUuId.String())
	return spaceId
}
