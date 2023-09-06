package applications

// CreateSpacesUseCase スペースの作成ユースケース
type CreateSpacesUseCase interface {

	// CreateSpace スペースの作成を実行する
	CreateSpace(input *CreateSpacesInput) (*CreateSpacesOutput, error)
}
