package attribute

import "errors"

type UserAttributeRecord struct {
	userAttributeName     UserAttributeName     // ユーザー属性名
	userAttributeType     UserAttributeType     // ユーザー属性値タイプ
	userAttributeValidate UserAttributeValidate // ユーザ属性バリデーション
	userAttributeValue    *UserAttributeValue   // ユーザー属性値
	canUpdate             bool                  // 変更可能フラグ
}

func NewUserAttributeRecord(
	userAttributeName UserAttributeName,
	userAttributeType UserAttributeType,
	userAttributeValue *UserAttributeValue,
	userAttributeValidate UserAttributeValidate, canUpdate bool) *UserAttributeRecord {

	return &UserAttributeRecord{
		userAttributeName:     userAttributeName,
		userAttributeType:     userAttributeType,
		userAttributeValue:    userAttributeValue,
		userAttributeValidate: userAttributeValidate,
		canUpdate:             canUpdate,
	}
}

// UpdateUserAttributeValue ユーザー属性値を変更する
func (a *UserAttributeRecord) UpdateUserAttributeValue(attrValue UserAttributeValue) error {
	// 変更拒否且つ、ユーザー属性値がnilでない場合は、変更できない。
	if !a.canUpdate && a.userAttributeValue != nil {
		return errors.New("ユーザー属性は、変更することができません。: " + a.userAttributeName.String())
	}

	a.userAttributeValue = &attrValue
	return nil
}

func (a *UserAttributeRecord) AttributeValue() *UserAttributeValue {
	return a.userAttributeValue
}

func (a *UserAttributeRecord) AttributeName() UserAttributeName {
	return a.userAttributeName
}

func (a *UserAttributeRecord) EqualTo(record UserAttributeRecord) bool {

	return a.userAttributeName.EqualTo(record.userAttributeName)
}
