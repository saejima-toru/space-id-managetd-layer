package users

import (
	"identity-management/domains/id-store/values/timestamp"
	"identity-management/domains/id-store/values/users"
	"identity-management/domains/id-store/values/users/status"
)

type User struct {
	userId         users.UserId             // ユーザーID
	userName       users.UserName           // ユーザー名
	userStatus     status.UserConfirmStatus // ユーザーステータス
	userAttributes UserAttributes           // ユーザー属性
	createAt       timestamp.TimeStamp      // ユーザー作成日
	updateAt       timestamp.TimeStamp      // ユーザー更新日
}

func NewUser(
	userId users.UserId,
	userName users.UserName,
	userStatus status.UserConfirmStatus,
	userAttributes UserAttributes,
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

func (u *User) UserId() users.UserId {
	return u.userId
}

func (u *User) UserName() users.UserName {
	return u.userName
}

func (u *User) UserStatus() status.UserConfirmStatus {
	return u.userStatus
}

func (u *User) UserAttributes() UserAttributes {
	return u.userAttributes
}

func (u *User) CreateAt() timestamp.TimeStamp {
	return u.createAt
}

func (u *User) UpdateAt() timestamp.TimeStamp {
	return u.updateAt
}
