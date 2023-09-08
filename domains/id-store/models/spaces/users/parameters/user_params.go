package parameters

import (
	"identity-management/domains/id-store/values/users"
)

/*
ユーザーパラメーター
*/

type UserParams struct {
	UserId                 *users.UserId           // 認証ユーザーID
	UserName               *users.UserName         // ユーザー名
	UserAttributesParams   *UserAttributesParam    // ユーザー属性値
	UserAuthenticateParams *UserAuthenticateParams // ユーザー認証情報
}
