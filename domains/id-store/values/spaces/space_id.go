package spaces

import "errors"

type SpaceId struct {
	userPoolId string
}

func NewUserPoolId(userPoolId string) (*SpaceId, error) {
	if len(userPoolId) < 1 {
		return nil, errors.New("ユーザープールIDの指定は、必須です。")
	}

	return &SpaceId{
		userPoolId: userPoolId,
	}, nil
}

// EqualTo 同一ユーザープールIDであるかを比較する
func (u *SpaceId) EqualTo(compareUserPoolId SpaceId) bool {
	return u.String() == compareUserPoolId.String()
}

func (u *SpaceId) String() string {
	return u.userPoolId
}
