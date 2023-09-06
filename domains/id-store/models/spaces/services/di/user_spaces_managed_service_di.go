package di

import (
	"identity-management/domains/id-store/models/spaces/services"
	"identity-management/domains/id-store/models/spaces/services/user_spaces"
	"identity-management/infrastructure/spaces/di"
)

func NewUserSpacesManagedServiceDI() services.SpacesManagedService {

	return user_spaces.NewUserSpacesManagedService(
		di.NewUserSpacesRepositoryForAws(),
	)
}
