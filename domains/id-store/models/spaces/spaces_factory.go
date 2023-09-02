package spaces

import "identity-management/domains/id-store/models/spaces/services/parameters"

// TODO: 実装

type CreateSpacesFactory interface {

	// CreateNewSpace 新規スペースを構成する
	CreateNewSpace(parameter parameters.SpacesParameter) (*Spaces, error)
}

type CreateUserSpacesFactory struct{}

func NewCreateUserSpacesFactory() *CreateUserSpacesFactory {
	return &CreateUserSpacesFactory{}
}

func (c *CreateUserSpacesFactory) CreateNewSpace(
	parameter parameters.SpacesParameter) (*Spaces, error) {

	//TODO implement me
	panic("implement me")
}
