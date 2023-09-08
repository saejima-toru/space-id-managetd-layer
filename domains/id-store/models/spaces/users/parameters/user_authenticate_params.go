package parameters

import "identity-management/domains/id-store/values/users"

/*
ユーザー認証パラメーター
*/

type UserAuthenticateParams struct {
	Password *users.Password
}
