package spaces

import (
	"errors"
	"regexp"
)

type NameSpace struct {
	nameSpace string
}

func NewNameSpace(nameSpace string) (*NameSpace, error) {
	if len(nameSpace) < 1 || len(nameSpace) > 128 {
		return nil, errors.New("ネームスペースは、1〜128文字である必要があります。")
	}
	if validate := regexp.MustCompile("^.*[^a-zA-Z\\d+=,.@-].*$"); validate.MatchString(nameSpace) {
		return nil, errors.New("ネームスペースは、大文字英字・番号(0〜9)・特殊文字(+ = , . @ -)のみ含めることができます。")
	}

	return &NameSpace{
		nameSpace: nameSpace,
	}, nil
}

func (p *NameSpace) String() string {
	return p.nameSpace
}
