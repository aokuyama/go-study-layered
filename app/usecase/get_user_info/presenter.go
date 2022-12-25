package get_user_info

import (
	"fmt"
	"net/http"

	"github.com/aokuyama/go-study-layered/app/usecase"
	"github.com/aokuyama/go-study-layered/app/view"
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

type GetUserInfoPresenterHttp struct {
	View view.HttpView
}

func NewGetUserInfoPresenterHttp(v view.HttpView) GetUserInfoPresenterHttp {
	return GetUserInfoPresenterHttp{v}
}

func (p GetUserInfoPresenterHttp) Render(o usecase.GetUserInfoOutput) error {
	if o.User == nil {
		p.View.String(http.StatusNotFound, "user not exist")
		return nil
	}
	v := usecase.GetUserInfoUserView{Name: o.User.FirstName.String()}
	p.View.Json(http.StatusOK, v)
	return nil
}
