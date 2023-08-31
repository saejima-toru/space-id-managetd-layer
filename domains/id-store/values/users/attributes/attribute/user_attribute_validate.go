package attribute

import "errors"

type UserAttributeValidate struct {
	minBytes int // 最小長
	maxBytes int // 最大長
}

func NewUserAttributeValidate(minBytes *int, maxBytes *int) (*UserAttributeValidate, error) {
	newValidate := &UserAttributeValidate{
		minBytes: 0,
		maxBytes: 2048,
	}

	if err := newValidate.setMinBytes(integer(minBytes, 0)); err != nil {
		return nil, err
	}
	if err := newValidate.setMaxBytes(integer(maxBytes, 2048)); err != nil {
		return nil, err
	}
	return newValidate, nil
}

func integer(v *int, defaults int) int {
	if v != nil {
		return *v
	}
	return defaults
}

func (a *UserAttributeValidate) setMaxBytes(maxBytes int) error {
	if maxBytes < 0 || maxBytes > 2048 {
		return errors.New("最大長の設定は、0〜2048バイト以下を指定する必要があります。")
	}
	a.maxBytes = maxBytes
	return nil
}
func (a *UserAttributeValidate) setMinBytes(minBytes int) error {
	if minBytes < 0 || minBytes > 2048 {
		return errors.New("最小長の設定は、0〜2048バイト以下を指定する必要があります。")
	}
	a.minBytes = minBytes
	return nil
}
