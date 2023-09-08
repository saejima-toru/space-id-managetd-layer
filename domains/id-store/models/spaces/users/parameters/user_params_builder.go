package parameters

import "identity-management/domains/id-store/values/users"

/**
ユーザーパラメータービルダー
*/

type userParamsBuilder struct {
	params *UserParams
}

func NewUserParamsBuilder() *userParamsBuilder {

	return &userParamsBuilder{
		params: &UserParams{},
	}
}

// UserName ユーザー名パラメータの設定
func (u *userParamsBuilder) UserName(userName *string) {
	if userName != nil {
		newUserName, _ := users.NewUserName(*userName)
		u.params.UserName = newUserName
	}
}

// UserId ユーザーIDパラメーターの設定
func (u *userParamsBuilder) UserId(userId *string) {
	if userId != nil {
		newUserId, _ := users.NewUserId(*userId)
		u.params.UserId = newUserId
	}
}
