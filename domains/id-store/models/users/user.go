package users

import (
	"identity-management/domains/id-store/values/timestamp"
	"identity-management/domains/id-store/values/users"
	"identity-management/domains/id-store/values/users/attributes"
	"identity-management/domains/id-store/values/users/status"
)

type User struct {
	userId         users.UserId              // ユーザーID
	userName       users.UserName            // ユーザー名
	userStatus     status.UserConfirmStatus  // ユーザーステータス
	userAttributes attributes.UserAttributes // ユーザー属性
	createAt       timestamp.TimeStamp       // ユーザー作成日
	updateAt       timestamp.TimeStamp       // ユーザー更新日
}

func NewUser(
	userId users.UserId,
	userName users.UserName,
	userStatus status.UserConfirmStatus,
	userAttributes attributes.UserAttributes,
	createAt timestamp.TimeStamp,
	updateAt timestamp.TimeStamp) *User {

	return &User{
		userId:         userId,
		userName:       userName,
		userStatus:     userStatus,
		userAttributes: userAttributes,
		createAt:       createAt,
		updateAt:       updateAt,
	}
}
