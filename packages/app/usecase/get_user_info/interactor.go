package get_user_info

import (
	"github.com/aokuyama/go-study-layered/packages/app/usecase"
	"github.com/aokuyama/go-study-layered/packages/domain/model/user"
)

type GetUserInfoInteractor struct {
	presenter  usecase.GetUserInfoPresenter
	repository user.UserRepository
}

func NewGetUserInfoInteractor(
	presenter usecase.GetUserInfoPresenter,
	repository user.UserRepository,
) GetUserInfoInteractor {
	return GetUserInfoInteractor{
		presenter,
		repository,
	}
}

func (u GetUserInfoInteractor) Handle(i usecase.GetUserInfoInput) error {
	user, err := u.repository.LoadByID(i.UserID)
	if err != nil {
		return err
	}
	o := usecase.GetUserInfoOutput{User: user}
	return u.presenter.Render(o)
}
