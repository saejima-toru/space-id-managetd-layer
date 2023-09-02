package spaces

/**
ID管理リポジトリ - 接続先
*/

import (
	"errors"
	"golang.org/x/exp/slices"
)

// ConnectIdr Connect Identity-Managed Repository
type ConnectIdr struct {
	connectIdr string
}

func NewConnectIdr(connectIdr string) (*ConnectIdr, error) {
	var idrs = []string{"aws_cognito"}
	if !slices.Contains(idrs, connectIdr) {
		return nil, errors.New("指定することができないリポジトリが選択されました。")
	}

	return &ConnectIdr{connectIdr: connectIdr}, nil
}

func (i *ConnectIdr) String() string {
	return i.connectIdr
}
