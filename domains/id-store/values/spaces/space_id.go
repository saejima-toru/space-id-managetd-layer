package spaces

import "errors"

type SpaceId struct {
	userPoolId string
}

func NewSpaceId(userPoolId string) (*SpaceId, error) {
	if len(userPoolId) < 1 {
		return nil, errors.New("スペース IDの指定は、必須です。")
	}

	return &SpaceId{
		userPoolId: userPoolId,
	}, nil
}

// EqualTo 同一スペースIDであるかを比較する
func (u *SpaceId) EqualTo(compareUserPoolId SpaceId) bool {
	return u.String() == compareUserPoolId.String()
}

func (u *SpaceId) String() string {
	return u.userPoolId
}
