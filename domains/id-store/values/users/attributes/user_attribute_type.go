package attributes

import "errors"

type UserAttributeType struct {
	userAttrType string
}

func NewUserAttributeType(userAttrType string) (*UserAttributeType, error) {
	var enums = []string{"string", "number"}

	if !contains(userAttrType, enums) {
		return nil, errors.New("ユーザー属性タイプに指定することができない値が指定されました。")
	}
	return &UserAttributeType{userAttrType: userAttrType}, nil
}

func contains(v string, slices []string) bool {
	for _, vType := range slices {
		if vType == v {
			return true
		}
	}
	return false
}
