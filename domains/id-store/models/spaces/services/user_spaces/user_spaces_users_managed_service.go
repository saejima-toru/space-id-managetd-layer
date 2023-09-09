package user_spaces

import (
	"identity-management/domains/id-store/models/spaces/services"
	"identity-management/domains/id-store/models/spaces/users"
	"identity-management/domains/id-store/models/spaces/users/parameters"
	"identity-management/domains/id-store/values/spaces"
	valueUser "identity-management/domains/id-store/values/users"
)

// UserSpacesUsersManagesService ユーザースペース：ユーザー管理サービス
type UserSpacesUsersManagesService struct {
	spacesUserRepository services.SpacesUserRepository
}

// CreateAuthenticateUser 認証ユーザーをユーザースペースに新規追加する
func (u *UserSpacesUsersManagesService) CreateAuthenticateUser(
	params parameters.UserParams) error {
	userFactory := users.NewCreateUserFactory()

	newUser, newUserErr := userFactory.CreateUser(params)
	if newUserErr != nil {
		return newUserErr
	}

	return u.spacesUserRepository.SaveUser(*newUser)
}

// DeleteUser ユーザースペース内、ユーザーの削除をする
func (u *UserSpacesUsersManagesService) DeleteUser(
	deleteUserId valueUser.UserId, spacesId spaces.SpaceId) error {
	//TODO implement me
	panic("implement me")
}

// DisabledAccessUser ユーザースペース内、ユーザーを無効化する
func (u *UserSpacesUsersManagesService) DisabledAccessUser(
	userId valueUser.UserId, id spaces.SpaceId) error {
	//TODO implement me
	panic("implement me")
}

// EnabledAccessUser ユーザースペース内、ユーザーを有効化する
func (u *UserSpacesUsersManagesService) EnabledAccessUser(
	userId valueUser.UserId, id spaces.SpaceId) error {
	//TODO implement me
	panic("implement me")
}

// ExistsUser ユーザースペース内に、ユーザーが存在するかを確認する
func (u *UserSpacesUsersManagesService) ExistsUser(
	existsCheckUser users.User, id spaces.SpaceId) bool {

	//TODO implement me
	panic("implement me")
}
