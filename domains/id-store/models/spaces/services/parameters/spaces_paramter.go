package parameters

import "identity-management/domains/id-store/values/spaces"

// SpacesParameter スペースパラメーター
type SpacesParameter struct {
	nameSpace *spaces.NameSpace
}

func (s *SpacesParameter) NameSpace() *spaces.NameSpace {
	return s.nameSpace
}

type SpaceParameterBuilder struct {
	params *SpacesParameter
}

func NewSpaceParameterBuilder() *SpaceParameterBuilder {
	return &SpaceParameterBuilder{
		params: new(SpacesParameter),
	}
}

// NameSpace スペースパラメータ、スペース名のセッター
func (s *SpaceParameterBuilder) NameSpace(namespace *string) {
	if namespace != nil {
		newNameSpace, _ := spaces.NewNameSpace(*namespace)
		s.params.nameSpace = newNameSpace
	}
}

func (s *SpaceParameterBuilder) Build() SpacesParameter {
	if s.params == nil {
		return *new(SpacesParameter)
	}
	return *s.params
}
