package status

import "errors"

func NewCantShitUserConfirmStatus() error {
	return errors.New("ユーザーステータスの移行を実行することができませんでした。")
}
