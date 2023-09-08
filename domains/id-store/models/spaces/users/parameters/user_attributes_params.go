package parameters

import (
	"identity-management/domains/id-store/values/users/attributes/attribute"
)

type UserAttributesParam map[string]UserAttributeValueParam
type UserAttributeValueParam struct {
	S *attribute.UserAttributeValue // 文字列
	N *attribute.UserAttributeValue // 数値
}
