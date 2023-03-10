package usecase

import "github.com/aokuyama/go-study-layered/packages/domain/model/user"

type GetUserInfoInput struct {
	UserID user.UserID
}

type GetUserInfoOutput struct {
	User *user.User
}

type GetUserInfoUseCase interface {
	Handle(i GetUserInfoInput) error
}

type GetUserInfoPresenter interface {
	Render(o GetUserInfoOutput) error
}

type GetUserInfoViewModel struct {
	IsFound bool
	Name    string
}

type GetUserInfoView interface {
	UpdateGetUserInfo(m GetUserInfoViewModel) error
}
