package services

import "identity-management/domains/id-store/models/spaces/users"

type SpacesUserRepository interface {

	// SaveUser ユーザー情報を保存する
	SaveUser(user users.User) error
}
