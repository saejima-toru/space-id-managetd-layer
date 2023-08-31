package attributes

// findByAttrName ユーザー属性名前検索
func (u *UserAttributes) findByAttrName(findAttrName UserAttributeName) *UserAttributeRecord {
	for _, record := range u.userAttributeRecords {
		attrName := record.userAttributeName
		if attrName.EqualTo(findAttrName) {
			return &record
		}
	}
	return nil
}

func (u *UserAttributes) indexOf(mixed UserAttributeRecord) int {
	for idx, record := range u.userAttributeRecords {
		if record.EqualTo(mixed) {
			return idx
		}
	}
	return -1
}

// contains ユーザー属性を含むか
func (u *UserAttributes) contains(mixed UserAttributeRecord) bool {
	for _, record := range u.userAttributeRecords {
		if record.EqualTo(mixed) {
			return true
		}
	}
	return false
}

// duplicate 重複を確認する
func duplicate(records []UserAttributeRecord) bool {
	encountered := map[string]bool{}
	for i := 0; i < len(records); i++ {
		attrName := records[i].userAttributeName
		if encountered[attrName.String()] {
			return true
		}
		encountered[attrName.String()] = true
	}
	return false
}
