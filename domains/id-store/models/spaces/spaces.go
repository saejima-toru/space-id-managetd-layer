package spaces

import (
	"identity-management/domains/id-store/values/spaces"
	"identity-management/domains/id-store/values/timestamp"
)

/**
ユーザープールの管理を担う。
*/

type Spaces interface{}

type Space struct {
	Spaces
	spaceId   spaces.SpaceId      // スペースID
	spaceName spaces.NameSpace    // ネームスペース
	createAt  timestamp.TimeStamp // ユーザープール作成日
	updateAt  timestamp.TimeStamp // ユーザープール更新日
}

func NewSpace(
	spaceId spaces.SpaceId,
	spaceName spaces.NameSpace,
	createAt timestamp.TimeStamp,
	updateAt timestamp.TimeStamp) *Space {

	return &Space{
		spaceId:   spaceId,
		spaceName: spaceName,
		createAt:  createAt,
		updateAt:  updateAt,
	}
}

func (u *Space) SpaceId() spaces.SpaceId {
	return u.spaceId
}

func (u *Space) SpaceName() spaces.NameSpace {
	return u.spaceName
}

func (u *Space) CreateAt() timestamp.TimeStamp {
	return u.createAt
}

func (u *Space) UpdateAt() timestamp.TimeStamp {
	return u.updateAt
}
