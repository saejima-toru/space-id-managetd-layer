package users

import (
	"errors"
	"github.com/google/uuid"
)

type UserId struct {
	userId string
}

func NewUserId(userId string) (*UserId, error) {
	if len(userId) < 1 {
		return nil, errors.New("ユーザーIDの指定は、必須です。")
	}
	return &UserId{
		userId: userId,
	}, nil
}
func NewGenerateUserId() UserId {
	uid, _ := uuid.NewUUID()

	return UserId{
		userId: uid.String(),
	}
}

func (u *UserId) EqualTo(userId UserId) bool {

	return u.userId == userId.userId
}
func (u *UserId) String() string {
	return u.userId
}
