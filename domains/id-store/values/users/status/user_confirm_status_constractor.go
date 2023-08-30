package status

// NewRegisteredUserConfirmStatus ユーザー登録済みステータス
func NewRegisteredUserConfirmStatus() *UserConfirmStatus {
	return silentNewUserConfirmStatus("registered")
}

// NewUserConfirmedStatus ユーザー確認済みステータス
func NewUserConfirmedStatus() *UserConfirmStatus {
	return silentNewUserConfirmStatus("confirmed")
}

// NewUserDisabledStatus ユーザーが無効されたステータス
func NewUserDisabledStatus() *UserConfirmStatus {
	return silentNewUserConfirmStatus("disabled")
}

// NewUserResetRequiredUserStatus ユーザーのパスワードリセット待ちステータス
func NewUserResetRequiredUserStatus() *UserConfirmStatus {
	return silentNewUserConfirmStatus("reset_required")
}

// NewForceChangePasswordStatus パスワード強制上書きステータス
func NewForceChangePasswordStatus() *UserConfirmStatus {
	return silentNewUserConfirmStatus("force_change_password")
}

func silentNewUserConfirmStatus(userConfirmStatus string) *UserConfirmStatus {
	status, _ := newUserConfirmStatus(userConfirmStatus)
	return status
}
