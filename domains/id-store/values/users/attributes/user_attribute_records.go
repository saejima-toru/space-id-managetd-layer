package attributes

import "errors"

type UserAttributeRecords struct {
	userAttributeRecords []UserAttributeRecord
}

func NewUserAttributeRecords(records []UserAttributeRecords) *UserAttributeRecords {

	return &UserAttributeRecords{
		userAttributeRecords: []UserAttributeRecord{},
	}
}

// AddUserAttributeRecord ユーザー属性を追加する
func (u *UserAttributeRecords) AddUserAttributeRecord(addAttrRecord UserAttributeRecord) error {
	if u.contains(addAttrRecord) {
		return errors.New("同じユーザー属性定義を追加することはできません。")
	}

	u.userAttributeRecords = append(u.userAttributeRecords, addAttrRecord)
	return nil
}

// UpdateUserAttributeValue 指定したユーザー属性値を変更する
func (u *UserAttributeRecords) UpdateUserAttributeValue(
	attrName UserAttributeName, attrValue UserAttributeValue) error {
	targetAttrRecord := u.findByAttrName(attrName)
	if targetAttrRecord != nil {
		return targetAttrRecord.UpdateUserAttributeValue(attrValue)
	}

	return nil
}

func (u *UserAttributeRecords) findByAttrName(findAttrName UserAttributeName) *UserAttributeRecord {
	for _, record := range u.userAttributeRecords {
		attrName := record.userAttributeName
		if attrName.EqualTo(findAttrName) {
			return &record
		}
	}
	return nil
}

func (u *UserAttributeRecords) indexOf(mixed UserAttributeRecord) int {
	for idx, record := range u.userAttributeRecords {
		if record.EqualTo(mixed) {
			return idx
		}
	}
	return -1
}

// contains ユーザー属性を含むか
func (u *UserAttributeRecords) contains(mixed UserAttributeRecord) bool {
	for _, record := range u.userAttributeRecords {
		if record.EqualTo(mixed) {
			return true
		}
	}
	return false
}
