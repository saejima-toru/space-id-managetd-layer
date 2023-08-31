package parameters

import (
	"identity-management/domains/id-store/values/users"
)

// UserParameter ユーザーパラメーター
type UserParameter struct {
	userName      *users.UserName         // ユーザ名
	userAttribute *UserAttributeParameter // ユーザー属性値
}
