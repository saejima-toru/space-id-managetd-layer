package parameters

import "identity-management/domains/id-store/values/spaces"

/*
スペースを検索するためのパラメーター
*/

type FetchSpacesParameter struct {
	spaceId   *spaces.SpaceId   // スペースID
	nameSpace *spaces.NameSpace // スペース名
}

// FetchSpaceParameterBuilder フェッチスペースパラメータービルダー
type FetchSpaceParameterBuilder struct {
	fetchSpacesParameter *FetchSpacesParameter
}

// SpaceId スペースIDを検索パラメーターに設定します。
func (f *FetchSpaceParameterBuilder) SpaceId(spaceId string) {
	spaceIdParam, err := spaces.NewSpaceId(spaceId)
	if err != nil {
		println("スキップ: スペースIDに指定できない値が指定されました。")
	}
	f.fetchSpacesParameter.spaceId = spaceIdParam
}

// NameSpace ネームスペースを検索パラメータに設定します。
func (f *FetchSpaceParameterBuilder) NameSpace(nameSpace string) {
	nameSpaceParam, err := spaces.NewNameSpace(nameSpace)
	if err != nil {
		println("スキップ: スペース名に指定できない値が指定されました。")
	}
	f.fetchSpacesParameter.nameSpace = nameSpaceParam
}

func NewFetchSpaceParameterBuilder() FetchSpaceParameterBuilder {

	return FetchSpaceParameterBuilder{
		fetchSpacesParameter: new(FetchSpacesParameter),
	}
}

func (f *FetchSpaceParameterBuilder) Build() FetchSpacesParameter {
	if f.fetchSpacesParameter == nil {
		println("FetchSpaceParameterBuilder: スペースの検索条件が指定されていません。")
		return FetchSpacesParameter{}
	}
	return *f.fetchSpacesParameter
}
