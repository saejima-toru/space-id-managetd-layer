package confirm_status

import "errors"

type UserConfirmStatus struct {
	userConfirmStatus string
}

func newUserConfirmStatus(userConfirmStatus string) (*UserConfirmStatus, error) {
	var types = []string{
		"registered", "confirmed", "disabled", "reset_required", "force_change_password"}
	if contains(userConfirmStatus, types) == false {
		return nil, errors.New("ユーザー確認ステータスに指定することができない値が指定されました。")
	}

	return &UserConfirmStatus{
		userConfirmStatus: userConfirmStatus,
	}, nil
}

func (s *UserConfirmStatus) EqualTo(status UserConfirmStatus) bool {
	return s.userConfirmStatus == status.userConfirmStatus
}

func contains(v string, types []string) bool {
	for _, vType := range types {
		if vType == v {
			return true
		}
	}
	return false
}
