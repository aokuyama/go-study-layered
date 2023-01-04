package get_user_info

import (
	"github.com/aokuyama/go-study-layered/packages/app/usecase"
)

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
