package confirm_status

// ShitToUserConfirmed ユーザー確認済みステータスへ移行する
func (s *UserConfirmStatus) ShitToUserConfirmed() (*UserConfirmStatus, error) {
	/*
		確認待ち、リセット対応待ち、強制パスワード書き換え時以外からは
		ステータス変更を行うことができない。
	*/
	if !s.IsUserRegisteredConfirmStatus() ||
		!s.IsUserResetRequiredStatus() || !s.IsUserForceChangePassword() {
		return nil, NewCantShitUserConfirmStatus()
	}

	return NewUserConfirmedStatus(), nil
}

// ShitToUserDisabledStatus ユーザーを無効化する
func (s *UserConfirmStatus) ShitToUserDisabledStatus() (*UserConfirmStatus, error) {
	if !s.IsUserConfirmedStatus() {
		return nil, NewCantShitUserConfirmStatus()
	}

	return NewUserDisabledStatus(), nil
}

// ShitToUserEnabledStatus ユーザーを有効ステータスへ移行する
func (s *UserConfirmStatus) ShitToUserEnabledStatus() (*UserConfirmStatus, error) {
	if !s.IsUserDisableStatus() {
		return nil, NewCantShitUserConfirmStatus()
	}

	return NewRegisteredUserConfirmStatus(), nil
}

// ShitToUserResetRequiredStatus ユーザーリセット待ちステータスへ移行する
func (s *UserConfirmStatus) ShitToUserResetRequiredStatus() (*UserConfirmStatus, error) {
	if !s.IsUserRegisteredConfirmStatus() {
		return nil, NewCantShitUserConfirmStatus()
	}

	return NewUserResetRequiredUserStatus(), nil
}
