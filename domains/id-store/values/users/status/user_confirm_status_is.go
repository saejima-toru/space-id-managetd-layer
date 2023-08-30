package status

// IsUserRegisteredConfirmStatus ユーザー未確認ステータス
func (s *UserConfirmStatus) IsUserRegisteredConfirmStatus() bool {

	return s.EqualTo(*NewUserResetRequiredUserStatus())
}

// IsUserConfirmedStatus ユーザー確認済みステータス
func (s *UserConfirmStatus) IsUserConfirmedStatus() bool {
	return s.EqualTo(*NewUserConfirmedStatus())
}

// IsUserDisableStatus ユーザー無効化ステータス
func (s *UserConfirmStatus) IsUserDisableStatus() bool {
	return s.EqualTo(*NewUserDisabledStatus())
}

// IsUserResetRequiredStatus ユーザーのパスワードリセット待機ステータス
func (s *UserConfirmStatus) IsUserResetRequiredStatus() bool {
	return s.EqualTo(*NewUserResetRequiredUserStatus())
}

// IsUserForceChangePassword ユーザーパスワードの強制変更済みステータス
func (s *UserConfirmStatus) IsUserForceChangePassword() bool {
	return s.EqualTo(*NewForceChangePasswordStatus())
}
