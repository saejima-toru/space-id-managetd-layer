package user

import (
	"identity-management/domains/id-store/values/timestamp"
	"identity-management/domains/id-store/values/users"
	"identity-management/domains/id-store/values/users/attributes"
	"identity-management/domains/id-store/values/users/status"
)

type User struct {
	userId         users.UserId
	userName       users.UserName
	userStatus     status.UserConfirmStatus
	userAttributes attributes.UserAttributes
	createAt       timestamp.TimeStamp
	updateAt       timestamp.TimeStamp
}
