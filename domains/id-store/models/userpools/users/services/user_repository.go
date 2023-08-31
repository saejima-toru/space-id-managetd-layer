package services

import (
	"identity-management/domains/id-store/models/userpools/users"
)

// UserRepository ユーザーリポジトリ
type UserRepository interface {

	// UpdateUser ユーザーを更新する
	UpdateUser(user users.User) error
}
