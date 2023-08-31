package attribute

type UserAttributeValue struct {
	userAttributeValue string
}

func NewUserAttributeValue(userAttributeValue string) *UserAttributeValue {

	return &UserAttributeValue{
		userAttributeValue: userAttributeValue,
	}
}

func (a *UserAttributeValue) String() string {
	return a.userAttributeValue
}
