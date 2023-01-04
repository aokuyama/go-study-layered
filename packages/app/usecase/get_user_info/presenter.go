package get_user_info

import (
	"fmt"

	"github.com/aokuyama/go-study-layered/packages/app/usecase"
)

type GetUserInfoPresenterCli struct {
}

func NewGetUserInfoPresenterCli() GetUserInfoPresenterCli {
	return GetUserInfoPresenterCli{}
}

func (p GetUserInfoPresenterCli) Render(o usecase.GetUserInfoOutput) error {
	if o.User != nil {
		fmt.Println(o.User.String())
	} else {
		fmt.Println("user not exist")
	}
	return nil
}

type GetUserInfoPresenter struct {
	View usecase.GetUserInfoView
}

func NewGetUserInfoPresenter(v usecase.GetUserInfoView) GetUserInfoPresenter {
	return GetUserInfoPresenter{v}
}

func (p GetUserInfoPresenter) Render(o usecase.GetUserInfoOutput) error {
	if o.User == nil {
		m := usecase.GetUserInfoViewModel{IsFound: false, Name: ""}
		return p.View.UpdateGetUserInfo(m)
	}
	m := usecase.GetUserInfoViewModel{IsFound: true, Name: o.User.FirstName.String()}
	return p.View.UpdateGetUserInfo(m)
}
