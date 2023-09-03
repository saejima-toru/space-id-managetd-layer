package parameters

import "identity-management/domains/id-store/values/spaces"

// SpacesParameter スペースパラメーター
type SpacesParameter struct {
	nameSpace spaces.NameSpace
}

func (s *SpacesParameter) NameSpace() *spaces.NameSpace {
	return &s.nameSpace
}
