package users

import "identity-management/domains/id-store/models/spaces/users/parameters"

// UserRepository ユーザーリポジトリ
type UserRepository interface {

	// CreateUser ユーザーの作成
	CreateUser(params *parameters.UserParams) (*User, error)
}
