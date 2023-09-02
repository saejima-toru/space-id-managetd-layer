package users

// UserRepository ユーザーリポジトリ
type UserRepository interface {

	// UpdateUser ユーザーを更新する
	UpdateUser(user User) error
}
