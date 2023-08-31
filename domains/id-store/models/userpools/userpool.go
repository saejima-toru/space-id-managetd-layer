package userpools

import (
	"identity-management/domains/id-store/values/timestamp"
	"identity-management/domains/id-store/values/userpools"
)

type UserPool struct {
	userPoolId   userpools.UserPoolId   // ユーザープールID
	userPoolName userpools.UserPoolName // ユーザープール名
	createAt     timestamp.TimeStamp    // ユーザープール作成日
	updateAt     timestamp.TimeStamp    // ユーザープール更新日
}
